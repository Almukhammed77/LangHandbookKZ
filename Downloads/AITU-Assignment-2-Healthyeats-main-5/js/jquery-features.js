
$(function () {
    console.log("jQuery is ready!");


    const $progress = $("#scrollProgress");
    const updateProgress = () => {
        const scrollTop = $(window).scrollTop();
        const docH = $(document).height() - $(window).height();
        const pct = docH > 0 ? (scrollTop / docH) * 100 : 0;
        $progress.css("width", pct + "%");
    };
    updateProgress();
    $(window).on("scroll resize", updateProgress);


    const $search = $("#searchBar");
    const $listItems = $("#recipeList .searchable");
    const $suggest = $("#searchSuggest");


    const SUGGESTIONS = $listItems
        .map(function () { return $(this).text().trim(); })
        .get();


    function highlightTerm($elements, term) {
        const t = term.trim();
        $elements.each(function () {
            const $el = $(this);
            const text = $el.text();
            if (!t) {
                $el.html(text); // снятие подсветки
                return;
            }
            // Экраним спецсимволы RegExp
            const safe = t.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
            const re = new RegExp(`(${safe})`, "gi");
            const html = text.replace(re, "<mark>$1</mark>");
            $el.html(html);
        });
    }


    function handleSearch() {
        const q = $search.val().toLowerCase();


        $listItems.each(function () {
            const $li = $(this);
            $li.toggle($li.text().toLowerCase().indexOf(q) > -1);
        });


        highlightTerm($listItems, q);


        $suggest.empty();
        if (!q) return;
        const matches = SUGGESTIONS.filter(s => s.toLowerCase().includes(q)).slice(0, 6);
        matches.forEach(m => {
            $suggest.append(`<li class="suggest-item" tabindex="0">${m}</li>`);
        });
    }

    $search.on("input keyup", handleSearch);


    $suggest.on("click keydown", ".suggest-item", function (e) {
        if (e.type === "click" || (e.type === "keydown" && (e.key === "Enter" || e.key === " "))) {
            const txt = $(this).text();
            $search.val(txt);
            handleSearch();
            $suggest.empty();
            // небольшой тост, что применён поиск
            toast(`Search applied: "${txt}"`);
        }
    });


    const animateCounter = ($el) => {
        const target = parseInt($el.attr("data-count"), 10) || 0;
        const duration = 1200; // ms
        const start = performance.now();

        function tick(now) {
            const p = Math.min(1, (now - start) / duration);
            const val = Math.floor(target * p);
            $el.text(val);
            if (p < 1) requestAnimationFrame(tick);
        }
        requestAnimationFrame(tick);
    };

    $(".stat-num").each(function () { animateCounter($(this)); });


    const $form = $("#demoForm");
    const $btn = $("#submitBtn");
    const $btnText = $btn.find(".btn-text");

    $form.on("submit", function (e) {
        e.preventDefault();
        if (!$form[0].checkValidity()) return;


        $btn.prop("disabled", true).addClass("loading");
        $btnText.text("Please wait...");


        setTimeout(() => {
            $btn.prop("disabled", false).removeClass("loading");
            $btnText.text("Submit");
            toast("Form submitted successfully");
            $form[0].reset();
        }, 1600);
    });


    const $toastWrap = $("#toastContainer");
    function toast(message = "Done", timeout = 2200) {
        const $t = $(
            `<div class="toast-item">
        <span>${message}</span>
      </div>`
        );
        $toastWrap.append($t);
        setTimeout(() => $t.addClass("show"));
        setTimeout(() => { $t.removeClass("show"); setTimeout(() => $t.remove(), 300); }, timeout);
    }

    $("#copyBtn").on("click", async function () {
        try {
            const text = $("#couponText").text().trim();
            await navigator.clipboard.writeText(text);
            $("#copyTooltip").text("Copied to clipboard! ✓");
            toast("Coupon copied");
        } catch {
            $("#copyTooltip").text("Copy failed");
        } finally {
            setTimeout(() => $("#copyTooltip").text(""), 1500);
        }
    });


    const $lazyImgs = $("img.lazy");
    const loadIfVisible = () => {
        const winTop = $(window).scrollTop();
        const winBot = winTop + $(window).height();
        $lazyImgs.each(function () {
            const $img = $(this);
            if ($img.data("loaded")) return;
            const top = $img.offset().top;
            if (top < winBot + 150) { // небольшой предзагрузочный буфер
                const realSrc = $img.attr("data-src");
                if (realSrc) {
                    $img.attr("src", realSrc);
                    $img.data("loaded", true).removeClass("lazy");
                }
            }
        });
    };
    loadIfVisible();
    $(window).on("scroll resize", loadIfVisible);
});
