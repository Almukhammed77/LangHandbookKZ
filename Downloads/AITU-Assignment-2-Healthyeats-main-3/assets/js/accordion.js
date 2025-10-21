document.addEventListener("DOMContentLoaded", function () {
  const faqHeaders = document.querySelectorAll(".faq-header");

  faqHeaders.forEach(header => {
    header.addEventListener("click", () => {
      const entry = header.parentElement;
      const body = entry.querySelector(".faq-body");

      document.querySelectorAll(".faq-entry").forEach(other => {
        if (other !== entry) {
          other.classList.remove("expanded");
          other.querySelector(".faq-body").style.maxHeight = null;
        }
      });

      entry.classList.toggle("expanded");
      if (entry.classList.contains("expanded")) {
        body.style.maxHeight = body.scrollHeight + "px";
      } else {
        body.style.maxHeight = null;
      }
    });
  });
});
