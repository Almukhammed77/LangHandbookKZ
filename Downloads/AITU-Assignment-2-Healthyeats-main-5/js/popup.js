document.addEventListener("DOMContentLoaded", function () {
  const popup = document.getElementById("popupForm");
  const openBtn = document.getElementById("openPopup");
  const closeBtn = document.getElementById("closePopup");

  if (!popup) return;

  popup.style.display = "none";

  if (openBtn) {
    openBtn.addEventListener("click", function (event) {
      event.preventDefault();
      popup.style.display = "flex";
    });
  }

  if (closeBtn) {
    closeBtn.addEventListener("click", function () {
      popup.style.display = "none";
    });
  }

  window.addEventListener("click", function (event) {
    if (event.target === popup) {
      popup.style.display = "none";
    }
  });
});
