

document.addEventListener("DOMContentLoaded", function () {
  const dateContainer = document.getElementById("dateTime");

  function updateTime() {
    const now = new Date();
    const options = { month: 'long', day: 'numeric', year: 'numeric', hour: 'numeric', minute: '2-digit', second: '2-digit' };
    dateContainer.textContent = now.toLocaleDateString("en-US", options);
  }

  if (dateContainer) {
    updateTime();
    setInterval(updateTime, 1000);
  }
});
