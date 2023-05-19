// primary source: https://stackoverflow.com/a/19099203
// licensed under: CC BY-SA 3.0

var url = document.getElementById('data-loc').innerHTML;

var pdfjsLib = window['pdfjs-dist/build/pdf'];

pdfjsLib.GlobalWorkerOptions.workerSrc = '/js/pdfjs/build/pdf.worker.js';

var currPage = 1; //Pages are 1-based not 0-based
var numPages = 0;
var thePDF = null;

var visiblePages = 5; // Number of visible pages
var firstVisiblePage = 1; // The index of the first visible page
var totalRenderedPages = 0; // The total number of rendered pages

var flags = {
	url: url,
	disableAutoFetch: true,
	disableStream: false
};

pdfjsLib.getDocument(flags).promise.then(function(pdf_) {
	pdf = pdf_;

	thePDF = pdf;

	numPages = pdf.numPages;

	pdf.getPage(1).then(handlePages);
});

function handlePages(page) {
    if (currPage === 1) {
        loc = 'first-page';
    } else {
        loc = 'mag-div';
    }

    var containerDiv = document.createElement('div');
    containerDiv.classList.add('center');
    document.getElementById(loc).appendChild(containerDiv);

    var loadingMessageDiv = document.createElement('div');
    loadingMessageDiv.classList.add('tmp-loading');
    loadingMessageDiv.innerHTML = 'PAGE LOADING';
    containerDiv.appendChild(loadingMessageDiv)

	var viewport = page.getViewport({ scale: 1 });

    var canvas = document.createElement('canvas');
    canvas.classList.add('pdf-viewer');
    canvas.classList.add(currPage % 2 === 0 ? 'align-right' : 'align-left')
    var context = canvas.getContext('2d');

    var desiredHeight = document.documentElement.clientHeight;
    var desiredWidth = document.documentElement.clientWidth;
    var scale;

    if (window.innerWidth >= 768) {
        var scaleY = desiredHeight / viewport.height;
        var scaleX = desiredWidth / viewport.width;
        scale = Math.min(scaleX, scaleY);
    } else {
        scale = 1.5;
    }

    viewport = page.getViewport({ scale: scale });
    var outputScale = Math.min(window.devicePixelRatio, 2) || 1;

    canvas.width = Math.floor(viewport.width * outputScale);
    canvas.height = Math.floor(viewport.height * outputScale);

	var renderContext = {
		canvasContext: context,
		viewport: viewport
	};

    if (outputScale !== 1) {
        context.setTransform(outputScale, 0, 0, outputScale, 0, 0);
    }

    var renderTask = page.render(renderContext);
    renderTask.promise.then(function () {
        loadingMessageDiv.remove();
	    containerDiv.appendChild(canvas);
    });

    currPage++;
    totalRenderedPages++;

    if (totalRenderedPages < visiblePages && currPage <= numPages) {
        thePDF.getPage(currPage).then(handlePages);
    }
    if (totalRenderedPages >= numPages) {
        document.getElementById('loading-message').innerHTML =
            'All pages are rendered!';
    }
}

function renderNextPage() {
    if (firstVisiblePage + visiblePages <= numPages) {
        firstVisiblePage++;
        thePDF.getPage(firstVisiblePage + visiblePages - 1)
            .then(function (page) {
                handlePages(page);
            });
    }
}

// TODO
function handleScroll() {
    var scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    var scrollHeight = document.documentElement.scrollHeight;
    var clientHeight = document.documentElement.clientHeight;
    //var scrollTrigger = 0.25;

    if (scrollTop + clientHeight >= scrollHeight) {
        renderNextPage();
    }
}

function debounce(func, wait) {
    var timeout;
    return function () {
        var context = this,
            args = arguments;
        var later = function () {
            timeout = null;
            func.apply(context, args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

window.addEventListener('scroll', debounce(handleScroll, 200));
