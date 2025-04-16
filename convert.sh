ROOT_DIR="/home/scottsmith/Projects/talks"

docker run \
  -e MARP_USER="$(id -u):$(id -g)" \
  --mount type=bind,source=$ROOT_DIR,target=/mnt/talks/ \
  marpteam/marp-cli \
  /mnt/talks/cognit-go/cognit-go.md \
  --allow-local-files \
  -o /mnt/talks/cognit-go/cognit-go.html

docker run \
  -e MARP_USER="$(id -u):$(id -g)" \
  --mount type=bind,source=$ROOT_DIR,target=/mnt/talks/ \
  marpteam/marp-cli \
  /mnt/talks/cognit-go/cognit-go.md \
  --allow-local-files \
  --pdf \
  -o /mnt/talks/cognit-go/cognit-go.pdf

docker run \
  -e MARP_USER="$(id -u):$(id -g)" \
  --mount type=bind,source=$ROOT_DIR,target=/mnt/talks/ \
  marpteam/marp-cli \
  /mnt/talks/dont-be-an-iter-hater/dont-be-an-iter-hater.md \
  --allow-local-files \
  -o /mnt/talks/dont-be-an-iter-hater/dont-be-an-iter-hater.html

docker run \
  -e MARP_USER="$(id -u):$(id -g)" \
  --mount type=bind,source=$ROOT_DIR,target=/mnt/talks/ \
  marpteam/marp-cli \
  /mnt/talks/dont-be-an-iter-hater/dont-be-an-iter-hater.md \
  --allow-local-files \
  --pdf \
  -o /mnt/talks/dont-be-an-iter-hater/dont-be-an-iter-hater.pdf
