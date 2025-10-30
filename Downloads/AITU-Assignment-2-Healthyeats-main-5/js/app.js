$(function () {
  console.log('jQuery ready');

  var $input = $('#liveSearch');
  if (!$input.length) { console.warn('No #liveSearch on page'); return; }

  var $cards = $('.menu .items .item');
  if (!$cards.length) $cards = $('.item');

  var titleSel = '.itemname';
  var $titles = $cards.find(titleSel);
  if (!$titles.length) { titleSel = 'h3'; $titles = $cards.find('h3'); }

  $titles.each(function(){ $(this).data('orig', $(this).text()); });

  function apply(q){
    q = (q || '').trim().toLowerCase();
    var any = false;

    $cards.each(function(){
      var $t = $(this).find(titleSel).first();
      var raw = $t.data('orig') || $t.text();
      var txt = raw.toLowerCase();
      var show = !q || txt.indexOf(q) > -1;
      $(this).toggle(show);
      if (show) any = true;

      if (!q) { $t.html(raw); }
      else {
        var re = new RegExp('(' + q.replace(/[.*+?^${}()|[\]\\]/g, '\\$&') + ')','ig');
        $t.html(raw.replace(re, '<mark>$1</mark>'));
      }
    });

    $('#noResults').toggle(!any);
  }

  $input.on('input keyup', function(){ apply(this.value); });

  $(document).on('click', '.search-btn', function(){
    apply($input.val());
  });

  $input.on('keydown', function(e){
    if (e.key === 'Enter') { e.preventDefault(); apply(this.value); }
  });

  apply('');
});