var bbut = document.getElementById("back")
var ibut = document.getElementById("info")
var dbut = document.getElementById("download")

bbut.addEventListener("click", function () {
	window.history.go(-1);
	return false;
});

ibut.addEventListener("click", function () {
	window.location = "/main/";
	return false;
});

dbut.addEventListener("click", function () {
    window.open(url, '_blank');
});
