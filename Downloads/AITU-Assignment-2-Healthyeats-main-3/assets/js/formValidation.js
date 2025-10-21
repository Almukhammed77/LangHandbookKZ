document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("contactForm");
  if (!form) return;

  // Создаём/находим бокс под ошибки
  let errorBox = form.querySelector(".form-error");
  if (!errorBox) {
    errorBox = document.createElement("div");
    errorBox.className = "form-error";
    form.prepend(errorBox);
  }

  const get = (sel) => document.getElementById(sel);
  const nameEl = get("name");
  const emailEl = get("email");
  const passEl = get("password");
  const confEl = get("confirm");

  // Функция подсветки полей
  const mark = (el, ok) => {
    if (!el) return;
    el.style.border = ok ? "2px solid #699635" : "2px solid #f2b200";
  };

  const validate = () => {
    const msgs = [];

    // Имя
    const nameOk = !!nameEl.value.trim();
    if (!nameOk) msgs.push("Please enter your name.");
    mark(nameEl, nameOk);

    // Email (гибкая проверка)
    const emailOk = /^[^\s@]+@[^\s@]+\.[a-zA-Z]{2,}$/.test(emailEl.value.trim());
    if (!emailOk) msgs.push("Please enter a valid email address.");
    mark(emailEl, emailOk);

    // Пароль: длина 8–20
    const len = passEl.value.length;
    const passLenOk = len >= 8 && len <= 20;
    if (!passLenOk) msgs.push("Password must be 8–20 characters long.");
    mark(passEl, passLenOk);

    // Совпадение паролей
    const matchOk = passEl.value === confEl.value && confEl.value.length > 0;
    if (!matchOk) msgs.push("Passwords do not match.");
    mark(confEl, matchOk);

    return msgs;
  };

  // Валидация на submit
  form.addEventListener("submit", (e) => {
    e.preventDefault();
    const msgs = validate();

    if (msgs.length) {
      errorBox.innerHTML = msgs.map(m => `<p>${m}</p>`).join("");
      errorBox.style.display = "block";
      errorBox.classList.remove("success");
      return;
    }

    errorBox.innerHTML = `<p class="success">Form submitted successfully!</p>`;
    errorBox.style.display = "block";
    [nameEl, emailEl, passEl, confEl].forEach(el => mark(el, true));
    form.reset();
    setTimeout(() => (errorBox.style.display = "none"), 2000);
  });

  [nameEl, emailEl, passEl, confEl].forEach(el => {
    el.addEventListener("input", () => {
      const msgs = validate();
      if (!msgs.length) {
        errorBox.style.display = "none";
      }
    });
  });
});
