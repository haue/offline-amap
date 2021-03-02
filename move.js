var maxzoom = 14;
var bounds = map.getBounds();
var currentZoom = ~~map.getZoom();
map.panTo(new AMap.LngLat(bounds.southWest.lng, bounds.northEast.lat));
map.setZoom(++currentZoom);
var xd = -1;
var stepX = window.innerWidth / 4;
var stepY = window.innerHeight / 4;
/*
->->->
<-<-<-
->->->
<-<-<-
*/
function move() {
  map.panBy(xd * stepX, null);
  var bs = map.getBounds();
  console.log(currentZoom, bs);
  if (bs.northEast.lng > bounds.northEast.lng) {
    console.log("←");
    xd = -xd;
    map.panBy(stepX, -stepY);
  }
  if (bs.southWest.lng < bounds.southWest.lng) {
    console.log("→");
    xd = -xd;
    map.panBy(-stepX, -stepY);
  }
  if (bs.southWest.lat < bounds.southWest.lat) {
    console.log("↓");
    xd = -1;
    map.panTo(new AMap.LngLat(bounds.southWest.lng, bounds.northEast.lat));
    map.setZoom(++currentZoom);
  }
}
var inv = setInterval(() => {
  move();
  if (currentZoom > maxzoom) {
    clearInterval(inv);
    console.log("----done");
  }
}, 1000);
