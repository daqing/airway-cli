bin:
  go build -o ./bin

install:
  go install github.com/daqing/airway-cli@latest

i: install
