package cmd

type notificationCfg struct {
	active bool
	img string
	title string
	content string
	// can add sound w/ go-toast
}

type Cfg struct {
	Notifications notificationCfg
	//Whatsapp whatsappCfg
	//Keep keepCfg
	//OneDrive onedriveCfg
	//Other output adapters..
}

