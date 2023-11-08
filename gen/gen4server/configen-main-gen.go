package gen4server

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p17ac1c06cc_impldaos_BookDaoImpl{})
    inst.register(&p17ac1c06cc_impldaos_CredentialDaoImpl{})
    inst.register(&p17ac1c06cc_impldaos_ExampleDaoImpl{})
    inst.register(&p17ac1c06cc_impldaos_PublicKeyDaoImpl{})
    inst.register(&p17ac1c06cc_impldaos_SecretKeyDaoImpl{})
    inst.register(&p17ac1c06cc_impldaos_TerminalDaoImpl{})
    inst.register(&p47eaaa4f28_implservices_BookServiceImpl{})
    inst.register(&p47eaaa4f28_implservices_CredentialServiceImpl{})
    inst.register(&p47eaaa4f28_implservices_ExampleServiceImpl{})
    inst.register(&p47eaaa4f28_implservices_LocalAgentImpl{})
    inst.register(&p47eaaa4f28_implservices_PublicKeyServiceImpl{})
    inst.register(&p47eaaa4f28_implservices_SecretKeyServiceImpl{})
    inst.register(&p47eaaa4f28_implservices_SignUpInAuthorizer{})
    inst.register(&p47eaaa4f28_implservices_TerminalServiceImpl{})
    inst.register(&p9dfa3ca364_data_ThisGroup{})
    inst.register(&pccd32ffa44_implconvertors_BookConvertorImpl{})
    inst.register(&pccd32ffa44_implconvertors_CredentialConvertorImpl{})
    inst.register(&pccd32ffa44_implconvertors_ExampleConvertorImpl{})
    inst.register(&pccd32ffa44_implconvertors_PublicKeyConvertorImpl{})
    inst.register(&pccd32ffa44_implconvertors_SecretKeyConvertorImpl{})
    inst.register(&pccd32ffa44_implconvertors_TerminalConvertorImpl{})
    inst.register(&ped7f874eb3_controllers_AuthExController{})
    inst.register(&ped7f874eb3_controllers_BookController{})
    inst.register(&ped7f874eb3_controllers_CredentialController{})
    inst.register(&ped7f874eb3_controllers_ExampleController{})
    inst.register(&ped7f874eb3_controllers_PublicKeyController{})
    inst.register(&ped7f874eb3_controllers_SecretKeyController{})
    inst.register(&ped7f874eb3_controllers_TerminalController{})


    return nil
}
