ROOT_DIR="/home/scottsmith/Projects/talks"

cd cognit-go
npx @marp-team/marp-cli@latest $ROOT_DIR/cognit-go/cognit-go.md \
  --allow-local-files \
  -o $ROOT_DIR//cognit-go/cognit-go.html

npx @marp-team/marp-cli@latest $ROOT_DIR/cognit-go/cognit-go.md \
  --allow-local-files \
  --pdf \
  -o $ROOT_DIR/cognit-go/cognit-go.pdf

npx @marp-team/marp-cli@latest $ROOT_DIR/dont-be-an-iter-hater/dont-be-an-iter-hater.md \
  --allow-local-files \
  -o $ROOT_DIR/dont-be-an-iter-hater/dont-be-an-iter-hater.html

npx @marp-team/marp-cli@latest $ROOT_DIR/dont-be-an-iter-hater/dont-be-an-iter-hater.md \
  --allow-local-files \
  --pdf \
  -o $ROOT_DIR/dont-be-an-iter-hater/dont-be-an-iter-hater.pdf
