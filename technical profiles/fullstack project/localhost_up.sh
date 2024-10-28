# Función para limpiar procesos en segundo plano
cleanup() {
    echo "Deteniendo procesos..."
    kill $backend_pid 2>/dev/null  # Detener el proceso del backend
    kill $tail_pid 2>/dev/null      # Detener el proceso de tail
    exit 0  # Cerrar el script
}

# Captura la señal SIGINT (Ctrl+C)
trap cleanup SIGINT

# Launch DB
sudo systemctl is-active --quiet mongodb.service || sudo systemctl start mongodb.service

# Launch BE
source ./BE/backend/bin/activate
python ./BE/djangoAPI/run_fastapi.py > backend.log 2>&1 &
backend_pid=$!
# Show BE outputs in current terminal
tail -f backend.log &
tail_pid=$!

# Navegar al directorio del frontend y lanzar Angular
cd ./FE/cocktail-app
npx ng serve --proxy-config proxy.conf.json
