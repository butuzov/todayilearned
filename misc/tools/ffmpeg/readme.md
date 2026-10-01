# FFMPEG

```shell
# Mass Convert mov to mp4 with downsample resize
for f in IMG_9195.MOV; do ffmpeg -i $f -vf "scale=trunc(iw/1.5/2)*2:trunc(ih/1.5/2)*2" \
 	-c:v libx264 -crf 22 -preset medium \
	-c:a aac -b:a 92k \
	-map_metadata 0 -map 0:v:0 -map '0:a?' "resized/${f/MOV/mp4}"; 
done
```