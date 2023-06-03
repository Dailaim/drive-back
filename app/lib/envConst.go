package lib

type WompiKey struct {
	Private   string
	Public    string
	Integrity string
}

func (w *WompiKey) SetWompiKey() {

	w.Private = CheckEnv("WOMPI_PRIVATE_KEY", "")
	w.Public = CheckEnv("WOMPI_PUBLIC_KEY", "")
	w.Integrity = CheckEnv("WOMPI_INTEGRITY_KEY", "")

	if w.Private == "" || w.Public == "" {
		panic("WOMPI_PRIVATE_KEY or WOMPI_PUBLIC_KEY not found")
	}

	if w.Integrity == "" {
		panic("WOMPI_INTEGRITY_KEY not found")
	}

}
