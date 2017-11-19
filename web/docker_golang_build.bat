SET _GO_VERSION=1.7.5
SET _IMAGE=maiastra/a-frame_web_server
SET _TAG=0.1

echo "remove the local image %_IMAGE%:%_TAG%"
docker rmi %_IMAGE%:%_TAG%


echo "build executable"
docker run --rm -v %cd%:/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:%_GO_VERSION% bash -c ./build.sh

echo "building new image"
docker build -t %_IMAGE%:%_TAG% .
echo "image build you can start with :"
echo "docker run -d -p 8888:8080 %_IMAGE%:%_TAG%"