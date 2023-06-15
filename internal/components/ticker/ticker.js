{{define "ticker_js"}}
let tickWidth
let duration
let unpublished = [{}];
let mainArticle = {};
let intervalSet = false;
let mainSelected

function isEmpty(obj) {
  return Object.keys(obj).length === 0;
}


let mTitle = document.getElementById("articleTitleMain")
let mAuth = document.getElementById("articleAuthorMain")
let mDate = document.getElementById("datePublishedMain")

window.onload = function() 
{
mainArticle.ArticleAuthor = `{{.Article.ArticleAuthor}}`;
mainArticle.ArticleTitle  = `{{.Article.ArticleTitle}}`;
mainArticle.ArticleText   = `{{.Article.ArticleText}}`;
mainArticle.DatePublished = `{{.Article.DatePublished}}`;
mainArticle.MainImg       = `{{.Article.MainImg}}`;
mainArticle.Index = -1;

  {{ range $k, $v := .Unpublished }}
  unpublished[{{$k}}] = {};
  unpublished[{{$k}}].ArticleAuthor = `{{ $v.ArticleAuthor}}`;
  unpublished[{{$k}}].ArticleTitle  = `{{ $v.ArticleTitle}}`;
  unpublished[{{$k}}].ArticleText   = `{{ $v.ArticleText }}`;
  unpublished[{{$k}}].DatePublished = `{{ $v.DatePublished}}`;
  unpublished[{{$k}}].MainImg       = `{{ $v.MainImg}}`;
  unpublished[{{$k}}].Index         = {{$k}};
  if ({{$k}} == 11) {
    intervalSet = true;
    setMainArticle();
    setInterval(function(){
      setMainArticle()
    }, 3000)
  }
  {{end}}
  function setMainArticle() {
    if (mainSelected != null) {
      mainSelected.style.backgroundColor = "white";
    }
    let mainImg = document.getElementById("mainImg");
    if (unpublished.length > 1) {
      if (isEmpty(mainArticle)) {
        // mainArticle = unpublished[0];
        // mainSelected = document.getElementById("articleItem_" + 0);
      } else {
        mainArticle = unpublished[mainArticle.Index+1];
        mTitle.innerHTML = mainArticle.ArticleTitle;
        mAuth.innerHTML  = mainArticle.ArticleAuthor;
        mDate.innerHTML = mainArticle.DatePublished;
        mainSelected = document.getElementById("articleItem_" + mainArticle.Index);
      }
      mainImg.style.backgroundImage = "url('" + mainArticle.MainImg + "')"
    } else {
      mainImg.style.backgroundImage = "url(" + '{{ .MainImg }}' + ")"
      mainSelected = document.getElementById("articleItem_" + 0);
    }
    // mainSelected.style.backgroundColor = "#fffce7";
  }


  if (!intervalSet) {
    setMainArticle();
  }

  window.onscroll = function()
  {
    // if (window.innerHeight > window.innerWidth) {
    //    if (window.scrollY < window.innerHeight) {
    //       homeDiv = document.getElementById("articleItem_0");
    //       homeDiv.style.opacity = (100 - (window.scrollY/6)) / 100;
    //    }
    // }
  }

  // tickWidth = document.getElementById("tickerOuter").offsetWidth / ( window.innerWidth)
  // duration = (tickWidth * 90)

  // document.getElementById("ticker").style.animation = null;
  // document.getElementById("ticker").offsetHeight;
  // document.getElementById("ticker").style.animation = "linear " + duration + "s scrollLeft";
  // document.getElementById("ticker").style.animationIterationCount = "infinite";
}

var observer = new IntersectionObserver(function (entries) {
  if (!entries[0].isIntersecting) {
    document.getElementById("ticker").style.animation = null;
    document.getElementById("ticker").offsetHeight;
    document.getElementById("ticker").style.animation = "linear 500s scrollLeft";
    document.getElementById("ticker").style.animationIterationCount = "infinite";
  }
});

observer.observe(document.querySelector("#ticker"))
{{end}}
