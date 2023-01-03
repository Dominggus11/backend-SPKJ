package controllers

func BeforeNormalisasi(UjianSekolah float32, RerataRaport float32, IPA float32, IPS float32, Minat string) (float32, float32, float32, float32, float32) {
	var (
		ci_ujian_sekolah float32
		ci_rerata        float32
		ci_ipa           float32
		ci_ips           float32
		ci_minat         float32
	)

	if UjianSekolah > 91 {
		ci_ujian_sekolah = 5
	} else if UjianSekolah > 80 {
		ci_ujian_sekolah = 4
	} else if UjianSekolah > 70 {
		ci_ujian_sekolah = 3
	} else if UjianSekolah > 60 {
		ci_ujian_sekolah = 2
	} else {
		ci_ujian_sekolah = 1
	}

	if RerataRaport > 91 {
		ci_rerata = 5
	} else if RerataRaport > 80 {
		ci_rerata = 4
	} else if RerataRaport > 70 {
		ci_rerata = 3
	} else if RerataRaport > 60 {
		ci_rerata = 2
	} else {
		ci_rerata = 1
	}

	if IPA > 91 {
		ci_ipa = 5
	} else if IPA > 80 {
		ci_ipa = 4
	} else if IPA > 70 {
		ci_ipa = 3
	} else if IPA > 60 {
		ci_ipa = 2
	} else {
		ci_ipa = 1
	}

	if IPS > 91 {
		ci_ips = 5
	} else if IPS > 80 {
		ci_ips = 4
	} else if IPS > 70 {
		ci_ips = 3
	} else if IPS > 60 {
		ci_ips = 2
	} else {
		ci_ips = 1
	}

	if Minat == "IPA" {
		ci_minat = 5
	} else {
		ci_minat = 1
	}

	return ci_ujian_sekolah, ci_rerata, ci_ipa, ci_ips, ci_minat
}

func AfterNormalisasi() {

}

func ResultJurusan() {

}
