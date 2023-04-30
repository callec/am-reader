var bbut = document.getElementById("back")
var ibut = document.getElementById("info")
var nbut = document.getElementById("next")

bbut.addEventListener("click", function () {
	window.history.go(-1);
	return false;
});

ibut.addEventListener("click", function () {
	window.location = "/main/";
	return false;
});

nbut.addEventListener("click", function () {
	// TODO
});
