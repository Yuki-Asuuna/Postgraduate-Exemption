go build
PID=$(ps -ef | grep Postgraduate-Exemption | grep -v grep| awk '{print $2}')
kill -9 $PID
./Postgraduate-Exemption