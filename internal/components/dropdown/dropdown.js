{{define "dropdown_js"}}
let menuItems = document.getElementById("menuItemOuter");
function toggleMenu() {
    let dd = document.getElementById("ddOuter");
    let nav = document.getElementById("nav")

    dd.style.animation = null;
    dd.offsetHeight;       
    if (menuItems.style.display != "none") {
        menuItems.style.display = "none";
        nav.style.display = "unset";
        // dd.style.animation = '0.4s linear fadeOut';
        // dd.style.animationFillMode = "forwards";
        // navName.style.color = '#fbbd00';
    } else {
        menuItems.style.display = "unset";
        nav.style.display = "none";
        // dd.style.animation = '0.2s linear fadeIn';
        // dd.style.animationFillMode = "forwards";
    }
}
{{end}}
