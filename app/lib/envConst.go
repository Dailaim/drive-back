package lib

type WompiKey struct {
	Private string
	Public  string
}

func (w *WompiKey) SetWompiKey() {

	w.Private = CheckEnv("WOMPI_PRIVATE_KEY", "")
	w.Public = CheckEnv("WOMPI_PUBLIC_KEY", "")

	if w.Private == "" || w.Public == "" {
		panic("WOMPI_PRIVATE_KEY or WOMPI_PUBLIC_KEY not found")
	}

}
