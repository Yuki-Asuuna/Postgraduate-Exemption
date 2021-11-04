go build -o ./output/
PID=$(ps -ef | grep Postgraduate-Exemption | grep -v grep| awk '{print $2}')
kill -9 $PID
./output/Postgraduate-Exemption