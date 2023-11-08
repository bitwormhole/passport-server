package implservices

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/publickeys"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/auth"
)

// SignUpInAuthorizer 实现[注册-登录]二合一的授权组件
type SignUpInAuthorizer struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

	UserDao          rbacdb.UserDAO         //starter:inject("#")
	EmailDao         rbacdb.EmailAddressDAO //starter:inject("#")
	UserConvertor    rbacdb.UserConvertor   //starter:inject("#")
	AuthService      auth.Service           //starter:inject("#")
	PublicKeyService publickeys.Service     //starter:inject("#")

}

func (inst *SignUpInAuthorizer) _impl() (auth.Authorizer, auth.Mechanism, auth.Registry) {
	return inst, inst, inst
}

// ListRegistrations ...
func (inst *SignUpInAuthorizer) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Enabled:       true,
		Priority:      0,
		Mechanism:     inst,
		Authorizer:    inst,
		Authenticator: nil,
	}
	return []*auth.Registration{r1}
}

// Support ...
func (inst *SignUpInAuthorizer) Support(r auth.Request) bool {
	return r.Action() == "sign-up-in"
}

// Authorize ...
func (inst *SignUpInAuthorizer) Authorize(a auth.Authorization) error {

	// get email address
	email, err := inst.getEmailIdentity(a)
	if err != nil {
		return err
	}

	// find user
	user, ready, err := inst.findUser(email)
	if user != nil && err == nil {
		if ready {
			return inst.doLogin(a, email, user)
		}
		return fmt.Errorf("the user is not ready")
	}

	// do sign-up
	err = inst.doSignUp(a)
	if err != nil {
		return err
	}

	// do sign-in
	return inst.doLogin(a, email, user)
}

func (inst *SignUpInAuthorizer) makeRandomPassword() lang.Base64 {
	buffer := make([]byte, 28)
	rand.Read(buffer)
	return lang.Base64FromBytes(buffer)
}

func (inst *SignUpInAuthorizer) doSignUp(a auth.Authorization) error {

	ctx := a.Context()
	action := auth.ActionSignUp
	password := inst.makeRandomPassword()

	ab := auth.AuthorizationBuilder{
		Action:     action,
		Context:    ctx,
		Attributes: a.Attributes(),
		Parameters: a.Parameters(),
		Identities: a.Identities(),
	}
	a2 := ab.Create()
	a2.Parameters().SetParam("password", password.String())
	return inst.AuthService.Authorize(a2)
}

func (inst *SignUpInAuthorizer) doLogin(a1 auth.Authorization, email auth.EmailIdentity, user *rbacdb.UserEntity) error {

	ctx := a1.Context()
	action := auth.ActionLogin

	if user == nil {
		user2, ready, err := inst.findUser(email)
		if err == nil && ready && user2 != nil {
			user = user2
		} else {
			if err == nil {
				err = fmt.Errorf("user not ready")
			}
			return err
		}
	}

	err := inst.savePublicKey(a1, user)
	if err != nil {
		return err
	}

	user2, err := inst.UserConvertor.ConvertE2D(ctx, user)
	if err != nil {
		return err
	}

	uid := auth.NewUserIdentity(email.By(), user2)
	ab := auth.AuthorizationBuilder{
		Action:     action,
		Context:    ctx,
		Attributes: a1.Attributes(),
		Parameters: a1.Parameters(),
		Identities: []auth.Identity{uid},
	}
	a2 := ab.Create()
	return inst.AuthService.Authorize(a2)
}

func (inst *SignUpInAuthorizer) savePublicKey(a1 auth.Authorization, user *rbacdb.UserEntity) error {

	ctx := a1.Context()
	ser := inst.PublicKeyService
	o1 := &publickeys.DTO{}

	b64str, err := a1.Parameters().GetParamRequired("public_key")
	if err != nil {
		return err
	}

	data, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, o1)
	if err != nil {
		return err
	}

	// find pk by sum
	sum := inst.computePublicKeyFingerprint(o1)
	exists, err := ser.ExistSum(ctx, sum)
	if err == nil && exists {
		return nil
	}

	// insert pk
	o1.Owner = user.ID
	_, err = ser.Insert(ctx, o1)
	return err
}

func (inst *SignUpInAuthorizer) computePublicKeyFingerprint(pk *publickeys.DTO) lang.Hex {
	data := pk.Data.Bytes()
	sum := sha256.Sum256(data)
	return lang.HexFromBytes(sum[:])
}

func (inst *SignUpInAuthorizer) getEmailIdentity(a auth.Authorization) (auth.EmailIdentity, error) {
	ids := a.Identities()
	for _, ident := range ids {
		emid, ok := ident.(auth.EmailIdentity)
		if ok {
			return emid, nil
		}
	}
	return nil, fmt.Errorf("no auth.EmailIdentity in the auth.Authorization")
}

func (inst *SignUpInAuthorizer) findUser(emid auth.EmailIdentity) (*rbacdb.UserEntity, bool, error) {

	const ready = true
	addr1 := emid.Address()
	addr2, err := inst.EmailDao.FindByAddress(nil, addr1)
	if err != nil {
		return nil, !ready, err
	}

	user, err := inst.UserDao.Find(nil, addr2.Owner)
	if err != nil {
		return nil, !ready, err
	}

	// check the user
	if user.Email == addr2.ID && user.Enabled {
		return user, ready, nil
	}

	return user, !ready, nil
}
