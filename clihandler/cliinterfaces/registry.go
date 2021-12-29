package cliinterfaces

type Registry interface {
	Login()
	Scan()
	GetImageInfo()
	Search()
	GetScanStatus()
}
