bin:
  cd airway && go build -o ../bin

install:
  go install github.com/daqing/airway-cli/airway@latest

i: install
