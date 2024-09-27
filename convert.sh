# docker run \
#   --mount type=bind,source=/home/scottsmith/Documents/permanent/talks,target=/mnt/talks/ \
#   marpteam/marp-cli /mnt/talks/cognit-go/cognit-go.md --pdf -o /mnt/talks/cognit-go/cognit-go.pdf

cd cognit-go
npx @marp-team/marp-cli@latest /home/scottsmith/Documents/permanent/talks/cognit-go/cognit-go.md \
  --allow-local-files \
  -o /home/scottsmith/Documents/permanent/talks/cognit-go/cognit-go.html

npx @marp-team/marp-cli@latest /home/scottsmith/Documents/permanent/talks/cognit-go/cognit-go.md \
  --allow-local-files \
  --pdf \
  -o /home/scottsmith/Documents/permanent/talks/cognit-go/cognit-go.pdf

xdg-open /home/scottsmith/Documents/permanent/talks/cognit-go/cognit-go.pdf
