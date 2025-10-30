document.addEventListener("DOMContentLoaded", function () {
  const btn = document.getElementById("changeColorBtn");

  const colors = ["#FFFFFF", "#FFF8E1", "#F2B200", "#699635", "#D7F5D0", "#F9E79F"];
  let index = 0;

  if (btn) {
    btn.addEventListener("click", function () {
      document.body.style.backgroundColor = colors[index];
      index = (index + 1) % colors.length;
    });
  }
});
