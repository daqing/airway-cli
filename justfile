bin:
  go build -o ./bin/airway ./cmd/airway

install:
  go install github.com/daqing/airway-cli/cmd/airway@v0.0.1

i: install
