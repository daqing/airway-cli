bin:
  go build ./cmd/airway -o ./bin

install:
  go install github.com/daqing/airway-cli/cmd/airway@latest

i: install
