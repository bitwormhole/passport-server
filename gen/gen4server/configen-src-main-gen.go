package gen4server
import (
    pccd32ffa4 "github.com/bitwormhole/passport-server/app/components/implconvertors"
    p17ac1c06c "github.com/bitwormhole/passport-server/app/components/impldaos"
    p47eaaa4f2 "github.com/bitwormhole/passport-server/app/components/implservices"
    p9dfa3ca36 "github.com/bitwormhole/passport-server/app/data"
    p7fb1f792e "github.com/bitwormhole/passport-server/app/data/books"
    p1c89d0cfc "github.com/bitwormhole/passport-server/app/data/credentials"
    pd66c1dfde "github.com/bitwormhole/passport-server/app/data/example"
    pfe6ab876e "github.com/bitwormhole/passport-server/app/data/passportdb"
    p579fa53c1 "github.com/bitwormhole/passport-server/app/data/publickeys"
    pe74e1d1d4 "github.com/bitwormhole/passport-server/app/data/secretkeys"
    p297e70940 "github.com/bitwormhole/passport-server/app/data/terminals"
    ped7f874eb "github.com/bitwormhole/passport-server/app/web/controllers"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    pf5d2c6fae "github.com/starter-go/security-gorm/rbacdb"
    p9d209f7c2 "github.com/starter-go/security/auth"
    p9621e8b71 "github.com/starter-go/security/random"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type pccd32ffa4.BookConvertorImpl in package:github.com/bitwormhole/passport-server/app/components/implconvertors
//
// id:com-ccd32ffa4408ad81-implconvertors-BookConvertorImpl
// class:
// alias:alias-7fb1f792eb6ddf4cb7de948d244cd03a-Convertor
// scope:singleton
//
type pccd32ffa44_implconvertors_BookConvertorImpl struct {
}

func (inst* pccd32ffa44_implconvertors_BookConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccd32ffa4408ad81-implconvertors-BookConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-7fb1f792eb6ddf4cb7de948d244cd03a-Convertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccd32ffa44_implconvertors_BookConvertorImpl) new() any {
    return &pccd32ffa4.BookConvertorImpl{}
}

func (inst* pccd32ffa44_implconvertors_BookConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccd32ffa4.BookConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type pccd32ffa4.CredentialConvertorImpl in package:github.com/bitwormhole/passport-server/app/components/implconvertors
//
// id:com-ccd32ffa4408ad81-implconvertors-CredentialConvertorImpl
// class:
// alias:alias-1c89d0cfc58e0fe059e35ed19920ff91-Convertor
// scope:singleton
//
type pccd32ffa44_implconvertors_CredentialConvertorImpl struct {
}

func (inst* pccd32ffa44_implconvertors_CredentialConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccd32ffa4408ad81-implconvertors-CredentialConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-1c89d0cfc58e0fe059e35ed19920ff91-Convertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccd32ffa44_implconvertors_CredentialConvertorImpl) new() any {
    return &pccd32ffa4.CredentialConvertorImpl{}
}

func (inst* pccd32ffa44_implconvertors_CredentialConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccd32ffa4.CredentialConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type pccd32ffa4.ExampleConvertorImpl in package:github.com/bitwormhole/passport-server/app/components/implconvertors
//
// id:com-ccd32ffa4408ad81-implconvertors-ExampleConvertorImpl
// class:
// alias:alias-d66c1dfde1be2a15e195b4a5f341a5bc-Convertor
// scope:singleton
//
type pccd32ffa44_implconvertors_ExampleConvertorImpl struct {
}

func (inst* pccd32ffa44_implconvertors_ExampleConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccd32ffa4408ad81-implconvertors-ExampleConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-d66c1dfde1be2a15e195b4a5f341a5bc-Convertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccd32ffa44_implconvertors_ExampleConvertorImpl) new() any {
    return &pccd32ffa4.ExampleConvertorImpl{}
}

func (inst* pccd32ffa44_implconvertors_ExampleConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccd32ffa4.ExampleConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type pccd32ffa4.PublicKeyConvertorImpl in package:github.com/bitwormhole/passport-server/app/components/implconvertors
//
// id:com-ccd32ffa4408ad81-implconvertors-PublicKeyConvertorImpl
// class:
// alias:alias-579fa53c16fdefe7716b49528d94c365-Convertor
// scope:singleton
//
type pccd32ffa44_implconvertors_PublicKeyConvertorImpl struct {
}

func (inst* pccd32ffa44_implconvertors_PublicKeyConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccd32ffa4408ad81-implconvertors-PublicKeyConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-579fa53c16fdefe7716b49528d94c365-Convertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccd32ffa44_implconvertors_PublicKeyConvertorImpl) new() any {
    return &pccd32ffa4.PublicKeyConvertorImpl{}
}

func (inst* pccd32ffa44_implconvertors_PublicKeyConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccd32ffa4.PublicKeyConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type pccd32ffa4.SecretKeyConvertorImpl in package:github.com/bitwormhole/passport-server/app/components/implconvertors
//
// id:com-ccd32ffa4408ad81-implconvertors-SecretKeyConvertorImpl
// class:
// alias:alias-e74e1d1d49e5d519bbd8bafe955c3139-Convertor
// scope:singleton
//
type pccd32ffa44_implconvertors_SecretKeyConvertorImpl struct {
}

func (inst* pccd32ffa44_implconvertors_SecretKeyConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccd32ffa4408ad81-implconvertors-SecretKeyConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-e74e1d1d49e5d519bbd8bafe955c3139-Convertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccd32ffa44_implconvertors_SecretKeyConvertorImpl) new() any {
    return &pccd32ffa4.SecretKeyConvertorImpl{}
}

func (inst* pccd32ffa44_implconvertors_SecretKeyConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccd32ffa4.SecretKeyConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type pccd32ffa4.TerminalConvertorImpl in package:github.com/bitwormhole/passport-server/app/components/implconvertors
//
// id:com-ccd32ffa4408ad81-implconvertors-TerminalConvertorImpl
// class:
// alias:alias-297e709400bed1a320d3d4e635850885-Convertor
// scope:singleton
//
type pccd32ffa44_implconvertors_TerminalConvertorImpl struct {
}

func (inst* pccd32ffa44_implconvertors_TerminalConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ccd32ffa4408ad81-implconvertors-TerminalConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-297e709400bed1a320d3d4e635850885-Convertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pccd32ffa44_implconvertors_TerminalConvertorImpl) new() any {
    return &pccd32ffa4.TerminalConvertorImpl{}
}

func (inst* pccd32ffa44_implconvertors_TerminalConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pccd32ffa4.TerminalConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type p17ac1c06c.BookDaoImpl in package:github.com/bitwormhole/passport-server/app/components/impldaos
//
// id:com-17ac1c06ccb95ce0-impldaos-BookDaoImpl
// class:
// alias:alias-7fb1f792eb6ddf4cb7de948d244cd03a-DAO
// scope:singleton
//
type p17ac1c06cc_impldaos_BookDaoImpl struct {
}

func (inst* p17ac1c06cc_impldaos_BookDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17ac1c06ccb95ce0-impldaos-BookDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-7fb1f792eb6ddf4cb7de948d244cd03a-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17ac1c06cc_impldaos_BookDaoImpl) new() any {
    return &p17ac1c06c.BookDaoImpl{}
}

func (inst* p17ac1c06cc_impldaos_BookDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17ac1c06c.BookDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17ac1c06cc_impldaos_BookDaoImpl) getAgent(ie application.InjectionExt)pfe6ab876e.LocalAgent{
    return ie.GetComponent("#alias-fe6ab876eef421db93aee722222edaf3-LocalAgent").(pfe6ab876e.LocalAgent)
}


func (inst*p17ac1c06cc_impldaos_BookDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p17ac1c06c.CredentialDaoImpl in package:github.com/bitwormhole/passport-server/app/components/impldaos
//
// id:com-17ac1c06ccb95ce0-impldaos-CredentialDaoImpl
// class:
// alias:alias-1c89d0cfc58e0fe059e35ed19920ff91-DAO
// scope:singleton
//
type p17ac1c06cc_impldaos_CredentialDaoImpl struct {
}

func (inst* p17ac1c06cc_impldaos_CredentialDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17ac1c06ccb95ce0-impldaos-CredentialDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-1c89d0cfc58e0fe059e35ed19920ff91-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17ac1c06cc_impldaos_CredentialDaoImpl) new() any {
    return &p17ac1c06c.CredentialDaoImpl{}
}

func (inst* p17ac1c06cc_impldaos_CredentialDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17ac1c06c.CredentialDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17ac1c06cc_impldaos_CredentialDaoImpl) getAgent(ie application.InjectionExt)pfe6ab876e.LocalAgent{
    return ie.GetComponent("#alias-fe6ab876eef421db93aee722222edaf3-LocalAgent").(pfe6ab876e.LocalAgent)
}


func (inst*p17ac1c06cc_impldaos_CredentialDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p17ac1c06c.ExampleDaoImpl in package:github.com/bitwormhole/passport-server/app/components/impldaos
//
// id:com-17ac1c06ccb95ce0-impldaos-ExampleDaoImpl
// class:
// alias:alias-d66c1dfde1be2a15e195b4a5f341a5bc-DAO
// scope:singleton
//
type p17ac1c06cc_impldaos_ExampleDaoImpl struct {
}

func (inst* p17ac1c06cc_impldaos_ExampleDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17ac1c06ccb95ce0-impldaos-ExampleDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-d66c1dfde1be2a15e195b4a5f341a5bc-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17ac1c06cc_impldaos_ExampleDaoImpl) new() any {
    return &p17ac1c06c.ExampleDaoImpl{}
}

func (inst* p17ac1c06cc_impldaos_ExampleDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17ac1c06c.ExampleDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17ac1c06cc_impldaos_ExampleDaoImpl) getAgent(ie application.InjectionExt)pfe6ab876e.LocalAgent{
    return ie.GetComponent("#alias-fe6ab876eef421db93aee722222edaf3-LocalAgent").(pfe6ab876e.LocalAgent)
}


func (inst*p17ac1c06cc_impldaos_ExampleDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p17ac1c06c.PublicKeyDaoImpl in package:github.com/bitwormhole/passport-server/app/components/impldaos
//
// id:com-17ac1c06ccb95ce0-impldaos-PublicKeyDaoImpl
// class:
// alias:alias-579fa53c16fdefe7716b49528d94c365-DAO
// scope:singleton
//
type p17ac1c06cc_impldaos_PublicKeyDaoImpl struct {
}

func (inst* p17ac1c06cc_impldaos_PublicKeyDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17ac1c06ccb95ce0-impldaos-PublicKeyDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-579fa53c16fdefe7716b49528d94c365-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17ac1c06cc_impldaos_PublicKeyDaoImpl) new() any {
    return &p17ac1c06c.PublicKeyDaoImpl{}
}

func (inst* p17ac1c06cc_impldaos_PublicKeyDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17ac1c06c.PublicKeyDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17ac1c06cc_impldaos_PublicKeyDaoImpl) getAgent(ie application.InjectionExt)pfe6ab876e.LocalAgent{
    return ie.GetComponent("#alias-fe6ab876eef421db93aee722222edaf3-LocalAgent").(pfe6ab876e.LocalAgent)
}


func (inst*p17ac1c06cc_impldaos_PublicKeyDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p17ac1c06c.SecretKeyDaoImpl in package:github.com/bitwormhole/passport-server/app/components/impldaos
//
// id:com-17ac1c06ccb95ce0-impldaos-SecretKeyDaoImpl
// class:
// alias:alias-e74e1d1d49e5d519bbd8bafe955c3139-DAO
// scope:singleton
//
type p17ac1c06cc_impldaos_SecretKeyDaoImpl struct {
}

func (inst* p17ac1c06cc_impldaos_SecretKeyDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17ac1c06ccb95ce0-impldaos-SecretKeyDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-e74e1d1d49e5d519bbd8bafe955c3139-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17ac1c06cc_impldaos_SecretKeyDaoImpl) new() any {
    return &p17ac1c06c.SecretKeyDaoImpl{}
}

func (inst* p17ac1c06cc_impldaos_SecretKeyDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17ac1c06c.SecretKeyDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17ac1c06cc_impldaos_SecretKeyDaoImpl) getAgent(ie application.InjectionExt)pfe6ab876e.LocalAgent{
    return ie.GetComponent("#alias-fe6ab876eef421db93aee722222edaf3-LocalAgent").(pfe6ab876e.LocalAgent)
}


func (inst*p17ac1c06cc_impldaos_SecretKeyDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p17ac1c06c.TerminalDaoImpl in package:github.com/bitwormhole/passport-server/app/components/impldaos
//
// id:com-17ac1c06ccb95ce0-impldaos-TerminalDaoImpl
// class:
// alias:alias-297e709400bed1a320d3d4e635850885-DAO
// scope:singleton
//
type p17ac1c06cc_impldaos_TerminalDaoImpl struct {
}

func (inst* p17ac1c06cc_impldaos_TerminalDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17ac1c06ccb95ce0-impldaos-TerminalDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-297e709400bed1a320d3d4e635850885-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17ac1c06cc_impldaos_TerminalDaoImpl) new() any {
    return &p17ac1c06c.TerminalDaoImpl{}
}

func (inst* p17ac1c06cc_impldaos_TerminalDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17ac1c06c.TerminalDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17ac1c06cc_impldaos_TerminalDaoImpl) getAgent(ie application.InjectionExt)pfe6ab876e.LocalAgent{
    return ie.GetComponent("#alias-fe6ab876eef421db93aee722222edaf3-LocalAgent").(pfe6ab876e.LocalAgent)
}


func (inst*p17ac1c06cc_impldaos_TerminalDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p47eaaa4f2.BookServiceImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-BookServiceImpl
// class:
// alias:alias-7fb1f792eb6ddf4cb7de948d244cd03a-Service
// scope:singleton
//
type p47eaaa4f28_implservices_BookServiceImpl struct {
}

func (inst* p47eaaa4f28_implservices_BookServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-BookServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-7fb1f792eb6ddf4cb7de948d244cd03a-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_BookServiceImpl) new() any {
    return &p47eaaa4f2.BookServiceImpl{}
}

func (inst* p47eaaa4f28_implservices_BookServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.BookServiceImpl)
	nop(ie, com)

	
    com.Convertor = inst.getConvertor(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_BookServiceImpl) getConvertor(ie application.InjectionExt)p7fb1f792e.Convertor{
    return ie.GetComponent("#alias-7fb1f792eb6ddf4cb7de948d244cd03a-Convertor").(p7fb1f792e.Convertor)
}


func (inst*p47eaaa4f28_implservices_BookServiceImpl) getDao(ie application.InjectionExt)p7fb1f792e.DAO{
    return ie.GetComponent("#alias-7fb1f792eb6ddf4cb7de948d244cd03a-DAO").(p7fb1f792e.DAO)
}



// type p47eaaa4f2.CredentialServiceImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-CredentialServiceImpl
// class:
// alias:alias-1c89d0cfc58e0fe059e35ed19920ff91-Service
// scope:singleton
//
type p47eaaa4f28_implservices_CredentialServiceImpl struct {
}

func (inst* p47eaaa4f28_implservices_CredentialServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-CredentialServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-1c89d0cfc58e0fe059e35ed19920ff91-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_CredentialServiceImpl) new() any {
    return &p47eaaa4f2.CredentialServiceImpl{}
}

func (inst* p47eaaa4f28_implservices_CredentialServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.CredentialServiceImpl)
	nop(ie, com)

	
    com.Convertor = inst.getConvertor(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_CredentialServiceImpl) getConvertor(ie application.InjectionExt)p1c89d0cfc.Convertor{
    return ie.GetComponent("#alias-1c89d0cfc58e0fe059e35ed19920ff91-Convertor").(p1c89d0cfc.Convertor)
}


func (inst*p47eaaa4f28_implservices_CredentialServiceImpl) getDao(ie application.InjectionExt)p1c89d0cfc.DAO{
    return ie.GetComponent("#alias-1c89d0cfc58e0fe059e35ed19920ff91-DAO").(p1c89d0cfc.DAO)
}



// type p47eaaa4f2.ExampleServiceImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-ExampleServiceImpl
// class:
// alias:alias-d66c1dfde1be2a15e195b4a5f341a5bc-Service
// scope:singleton
//
type p47eaaa4f28_implservices_ExampleServiceImpl struct {
}

func (inst* p47eaaa4f28_implservices_ExampleServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-ExampleServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-d66c1dfde1be2a15e195b4a5f341a5bc-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_ExampleServiceImpl) new() any {
    return &p47eaaa4f2.ExampleServiceImpl{}
}

func (inst* p47eaaa4f28_implservices_ExampleServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.ExampleServiceImpl)
	nop(ie, com)

	
    com.Convertor = inst.getConvertor(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_ExampleServiceImpl) getConvertor(ie application.InjectionExt)pd66c1dfde.Convertor{
    return ie.GetComponent("#alias-d66c1dfde1be2a15e195b4a5f341a5bc-Convertor").(pd66c1dfde.Convertor)
}


func (inst*p47eaaa4f28_implservices_ExampleServiceImpl) getDao(ie application.InjectionExt)pd66c1dfde.DAO{
    return ie.GetComponent("#alias-d66c1dfde1be2a15e195b4a5f341a5bc-DAO").(pd66c1dfde.DAO)
}



// type p47eaaa4f2.LocalAgentImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-LocalAgentImpl
// class:
// alias:alias-fe6ab876eef421db93aee722222edaf3-LocalAgent
// scope:singleton
//
type p47eaaa4f28_implservices_LocalAgentImpl struct {
}

func (inst* p47eaaa4f28_implservices_LocalAgentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-LocalAgentImpl"
	r.Classes = ""
	r.Aliases = "alias-fe6ab876eef421db93aee722222edaf3-LocalAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_LocalAgentImpl) new() any {
    return &p47eaaa4f2.LocalAgentImpl{}
}

func (inst* p47eaaa4f28_implservices_LocalAgentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.LocalAgentImpl)
	nop(ie, com)

	
    com.DataSourceMan = inst.getDataSourceMan(ie)
    com.DataSourceAlias = inst.getDataSourceAlias(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_LocalAgentImpl) getDataSourceMan(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p47eaaa4f28_implservices_LocalAgentImpl) getDataSourceAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passport.datasource}")
}



// type p47eaaa4f2.PublicKeyServiceImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-PublicKeyServiceImpl
// class:
// alias:alias-579fa53c16fdefe7716b49528d94c365-Service
// scope:singleton
//
type p47eaaa4f28_implservices_PublicKeyServiceImpl struct {
}

func (inst* p47eaaa4f28_implservices_PublicKeyServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-PublicKeyServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-579fa53c16fdefe7716b49528d94c365-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_PublicKeyServiceImpl) new() any {
    return &p47eaaa4f2.PublicKeyServiceImpl{}
}

func (inst* p47eaaa4f28_implservices_PublicKeyServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.PublicKeyServiceImpl)
	nop(ie, com)

	
    com.Convertor = inst.getConvertor(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_PublicKeyServiceImpl) getConvertor(ie application.InjectionExt)p579fa53c1.Convertor{
    return ie.GetComponent("#alias-579fa53c16fdefe7716b49528d94c365-Convertor").(p579fa53c1.Convertor)
}


func (inst*p47eaaa4f28_implservices_PublicKeyServiceImpl) getDao(ie application.InjectionExt)p579fa53c1.DAO{
    return ie.GetComponent("#alias-579fa53c16fdefe7716b49528d94c365-DAO").(p579fa53c1.DAO)
}



// type p47eaaa4f2.SecretKeyServiceImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-SecretKeyServiceImpl
// class:
// alias:alias-e74e1d1d49e5d519bbd8bafe955c3139-Service
// scope:singleton
//
type p47eaaa4f28_implservices_SecretKeyServiceImpl struct {
}

func (inst* p47eaaa4f28_implservices_SecretKeyServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-SecretKeyServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-e74e1d1d49e5d519bbd8bafe955c3139-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_SecretKeyServiceImpl) new() any {
    return &p47eaaa4f2.SecretKeyServiceImpl{}
}

func (inst* p47eaaa4f28_implservices_SecretKeyServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.SecretKeyServiceImpl)
	nop(ie, com)

	
    com.Convertor = inst.getConvertor(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_SecretKeyServiceImpl) getConvertor(ie application.InjectionExt)pe74e1d1d4.Convertor{
    return ie.GetComponent("#alias-e74e1d1d49e5d519bbd8bafe955c3139-Convertor").(pe74e1d1d4.Convertor)
}


func (inst*p47eaaa4f28_implservices_SecretKeyServiceImpl) getDao(ie application.InjectionExt)pe74e1d1d4.DAO{
    return ie.GetComponent("#alias-e74e1d1d49e5d519bbd8bafe955c3139-DAO").(pe74e1d1d4.DAO)
}



// type p47eaaa4f2.SignUpInAuthorizer in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-SignUpInAuthorizer
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p47eaaa4f28_implservices_SignUpInAuthorizer struct {
}

func (inst* p47eaaa4f28_implservices_SignUpInAuthorizer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-SignUpInAuthorizer"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_SignUpInAuthorizer) new() any {
    return &p47eaaa4f2.SignUpInAuthorizer{}
}

func (inst* p47eaaa4f28_implservices_SignUpInAuthorizer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.SignUpInAuthorizer)
	nop(ie, com)

	
    com.UserDao = inst.getUserDao(ie)
    com.EmailDao = inst.getEmailDao(ie)
    com.UserConvertor = inst.getUserConvertor(ie)
    com.AuthService = inst.getAuthService(ie)
    com.PublicKeyService = inst.getPublicKeyService(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_SignUpInAuthorizer) getUserDao(ie application.InjectionExt)pf5d2c6fae.UserDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO").(pf5d2c6fae.UserDAO)
}


func (inst*p47eaaa4f28_implservices_SignUpInAuthorizer) getEmailDao(ie application.InjectionExt)pf5d2c6fae.EmailAddressDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-EmailAddressDAO").(pf5d2c6fae.EmailAddressDAO)
}


func (inst*p47eaaa4f28_implservices_SignUpInAuthorizer) getUserConvertor(ie application.InjectionExt)pf5d2c6fae.UserConvertor{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserConvertor").(pf5d2c6fae.UserConvertor)
}


func (inst*p47eaaa4f28_implservices_SignUpInAuthorizer) getAuthService(ie application.InjectionExt)p9d209f7c2.Service{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-Service").(p9d209f7c2.Service)
}


func (inst*p47eaaa4f28_implservices_SignUpInAuthorizer) getPublicKeyService(ie application.InjectionExt)p579fa53c1.Service{
    return ie.GetComponent("#alias-579fa53c16fdefe7716b49528d94c365-Service").(p579fa53c1.Service)
}



// type p47eaaa4f2.TerminalServiceImpl in package:github.com/bitwormhole/passport-server/app/components/implservices
//
// id:com-47eaaa4f289507e5-implservices-TerminalServiceImpl
// class:
// alias:alias-297e709400bed1a320d3d4e635850885-Service
// scope:singleton
//
type p47eaaa4f28_implservices_TerminalServiceImpl struct {
}

func (inst* p47eaaa4f28_implservices_TerminalServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-47eaaa4f289507e5-implservices-TerminalServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-297e709400bed1a320d3d4e635850885-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p47eaaa4f28_implservices_TerminalServiceImpl) new() any {
    return &p47eaaa4f2.TerminalServiceImpl{}
}

func (inst* p47eaaa4f28_implservices_TerminalServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p47eaaa4f2.TerminalServiceImpl)
	nop(ie, com)

	
    com.Convertor = inst.getConvertor(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p47eaaa4f28_implservices_TerminalServiceImpl) getConvertor(ie application.InjectionExt)p297e70940.Convertor{
    return ie.GetComponent("#alias-297e709400bed1a320d3d4e635850885-Convertor").(p297e70940.Convertor)
}


func (inst*p47eaaa4f28_implservices_TerminalServiceImpl) getDao(ie application.InjectionExt)p297e70940.DAO{
    return ie.GetComponent("#alias-297e709400bed1a320d3d4e635850885-DAO").(p297e70940.DAO)
}



// type p9dfa3ca36.ThisGroup in package:github.com/bitwormhole/passport-server/app/data
//
// id:com-9dfa3ca3646deb45-data-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p9dfa3ca364_data_ThisGroup struct {
}

func (inst* p9dfa3ca364_data_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9dfa3ca3646deb45-data-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9dfa3ca364_data_ThisGroup) new() any {
    return &p9dfa3ca36.ThisGroup{}
}

func (inst* p9dfa3ca364_data_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9dfa3ca36.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)


    return nil
}


func (inst*p9dfa3ca364_data_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.passport.enabled}")
}


func (inst*p9dfa3ca364_data_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passport.alias}")
}


func (inst*p9dfa3ca364_data_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passport.uri}")
}


func (inst*p9dfa3ca364_data_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passport.table-name-prefix}")
}


func (inst*p9dfa3ca364_data_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passport.datasource}")
}



// type ped7f874eb.AuthExController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-AuthExController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_AuthExController struct {
}

func (inst* ped7f874eb3_controllers_AuthExController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-AuthExController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_AuthExController) new() any {
    return &ped7f874eb.AuthExController{}
}

func (inst* ped7f874eb3_controllers_AuthExController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.AuthExController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service2 = inst.getService2(ie)
    com.Service1 = inst.getService1(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_AuthExController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_AuthExController) getService2(ie application.InjectionExt)p9d209f7c2.Service{
    return ie.GetComponent("#alias-9d209f7c2504d33e6054a2c9998e9485-Service").(p9d209f7c2.Service)
}


func (inst*ped7f874eb3_controllers_AuthExController) getService1(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}



// type ped7f874eb.BookController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-BookController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_BookController struct {
}

func (inst* ped7f874eb3_controllers_BookController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-BookController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_BookController) new() any {
    return &ped7f874eb.BookController{}
}

func (inst* ped7f874eb3_controllers_BookController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.BookController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_BookController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_BookController) getService(ie application.InjectionExt)p7fb1f792e.Service{
    return ie.GetComponent("#alias-7fb1f792eb6ddf4cb7de948d244cd03a-Service").(p7fb1f792e.Service)
}



// type ped7f874eb.CredentialController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-CredentialController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_CredentialController struct {
}

func (inst* ped7f874eb3_controllers_CredentialController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-CredentialController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_CredentialController) new() any {
    return &ped7f874eb.CredentialController{}
}

func (inst* ped7f874eb3_controllers_CredentialController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.CredentialController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_CredentialController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_CredentialController) getService(ie application.InjectionExt)p1c89d0cfc.Service{
    return ie.GetComponent("#alias-1c89d0cfc58e0fe059e35ed19920ff91-Service").(p1c89d0cfc.Service)
}



// type ped7f874eb.ExampleController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_ExampleController struct {
}

func (inst* ped7f874eb3_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_ExampleController) new() any {
    return &ped7f874eb.ExampleController{}
}

func (inst* ped7f874eb3_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_ExampleController) getService(ie application.InjectionExt)p7fb1f792e.Service{
    return ie.GetComponent("#alias-7fb1f792eb6ddf4cb7de948d244cd03a-Service").(p7fb1f792e.Service)
}



// type ped7f874eb.PublicKeyController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-PublicKeyController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_PublicKeyController struct {
}

func (inst* ped7f874eb3_controllers_PublicKeyController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-PublicKeyController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_PublicKeyController) new() any {
    return &ped7f874eb.PublicKeyController{}
}

func (inst* ped7f874eb3_controllers_PublicKeyController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.PublicKeyController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_PublicKeyController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_PublicKeyController) getService(ie application.InjectionExt)p579fa53c1.Service{
    return ie.GetComponent("#alias-579fa53c16fdefe7716b49528d94c365-Service").(p579fa53c1.Service)
}



// type ped7f874eb.SecretKeyController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-SecretKeyController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_SecretKeyController struct {
}

func (inst* ped7f874eb3_controllers_SecretKeyController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-SecretKeyController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_SecretKeyController) new() any {
    return &ped7f874eb.SecretKeyController{}
}

func (inst* ped7f874eb3_controllers_SecretKeyController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.SecretKeyController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_SecretKeyController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_SecretKeyController) getService(ie application.InjectionExt)pe74e1d1d4.Service{
    return ie.GetComponent("#alias-e74e1d1d49e5d519bbd8bafe955c3139-Service").(pe74e1d1d4.Service)
}



// type ped7f874eb.TerminalController in package:github.com/bitwormhole/passport-server/app/web/controllers
//
// id:com-ed7f874eb33e9777-controllers-TerminalController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type ped7f874eb3_controllers_TerminalController struct {
}

func (inst* ped7f874eb3_controllers_TerminalController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed7f874eb33e9777-controllers-TerminalController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped7f874eb3_controllers_TerminalController) new() any {
    return &ped7f874eb.TerminalController{}
}

func (inst* ped7f874eb3_controllers_TerminalController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped7f874eb.TerminalController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*ped7f874eb3_controllers_TerminalController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*ped7f874eb3_controllers_TerminalController) getService(ie application.InjectionExt)p297e70940.Service{
    return ie.GetComponent("#alias-297e709400bed1a320d3d4e635850885-Service").(p297e70940.Service)
}


