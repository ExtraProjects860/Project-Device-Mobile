package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var notAccepted = []string{
	"00000000000", "11111111111", "22222222222",
	"33333333333", "44444444444", "55555555555",
	"66666666666", "77777777777", "88888888888",
	"99999999999",
}

// GenerateCPF generates a random, valid CPF number in unformatted form (11 digits).
func GenerateCPF() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder

	for i := 0; i < 9; i++ {
		sb.WriteByte(byte('0' + rng.Intn(10)))
	}

	cpfBase := sb.String()
	dv1 := calculateFirstDigit(cpfBase)
	dv2 := calculateSecondDigit(cpfBase + strconv.Itoa(dv1))

	return fmt.Sprintf("%s%d%d", cpfBase, dv1, dv2)
}

// FormatCPF takes a CPF string (with or without formatting) and returns it in the formatted style: XXX.XXX.XXX-XX
func FormatCPF(cpf string) string {
	cpf = UnformatCPF(cpf)

	if len(cpf) != 11 {
		return "Invalid CPF"
	}

	mask := "###.###.###-##"
	result := make([]rune, len(mask))
	cpfIndex := 0

	for i, r := range mask {
		if r == '#' {
			result[i] = rune(cpf[cpfIndex])
			cpfIndex++
		} else {
			result[i] = r
		}
	}
	return string(result)
}

// UnformatCPF removes all non-numeric characters from the CPF string.
func UnformatCPF(cpf string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(cpf, "")
}

// ValidateCPF verifies if a given CPF is syntactically valid.
func ValidateCPF(value string) bool {
	cpf := UnformatCPF(value)
	if len(cpf) != 11 || inNotAccepted(cpf) {
		return false
	}

	dv1 := calculateFirstDigit(cpf[:9])
	dv2 := calculateSecondDigit(cpf[:9] + strconv.Itoa(dv1))

	return cpf[9] == byte('0'+dv1) && cpf[10] == byte('0'+dv2)
}

// Origin returns the Brazilian region associated with the CPF based on the 9th digit.
func Origin(value string) string {
	cpf := UnformatCPF(value)
	if len(cpf) != 11 {
		return ""
	}
	regionDigit, err := strconv.Atoi(string(cpf[8]))
	if err != nil {
		return ""
	}
	switch regionDigit {
	case 0:
		return "Rio Grande do Sul"
	case 1:
		return "Distrito Federal, Goiás, Mato Grosso do Sul, and Tocantins"
	case 2:
		return "Pará, Amazonas, Acre, Amapá, Rondônia, and Roraima"
	case 3:
		return "Ceará, Maranhão, and Piauí"
	case 4:
		return "Pernambuco, Rio Grande do Norte, Paraíba, and Alagoas"
	case 5:
		return "Bahia and Sergipe"
	case 6:
		return "Minas Gerais"
	case 7:
		return "Rio de Janeiro and Espírito Santo"
	case 8:
		return "São Paulo"
	case 9:
		return "Paraná and Santa Catarina"
	default:
		return ""
	}
}

// inNotAccepted checks if the CPF is one of the known invalid patterns (e.g., all digits the same).
func inNotAccepted(cpf string) bool {
	for _, invalid := range notAccepted {
		if cpf == invalid {
			return true
		}
	}
	return false
}

// calculateFirstDigit computes the first verification digit from the CPF base.
func calculateFirstDigit(base string) int {
	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(base[i] - '0')
		sum += digit * (10 - i)
	}
	rest := (sum * 10) % 11
	if rest == 10 {
		return 0
	}
	return rest
}

// calculateSecondDigit computes the second verification digit using the base + first digit.
func calculateSecondDigit(base string) int {
	sum := 0
	for i := 0; i < 10; i++ {
		digit := int(base[i] - '0')
		sum += digit * (11 - i)
	}
	rest := (sum * 10) % 11
	if rest == 10 {
		return 0
	}
	return rest
}
