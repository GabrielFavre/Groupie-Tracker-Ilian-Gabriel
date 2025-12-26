document.addEventListener('DOMContentLoaded', () => {
    const toggleButton = document.getElementById('theme-toggle');
    const body = document.body;
    const searchInput = document.getElementById('searchBar');
    const cards = document.querySelectorAll('.ds-cartridge');

    if (localStorage.getItem('theme') === 'dark') {
        body.classList.add('dark-mode');
        if (toggleButton) toggleButton.textContent = '☀️';
    }

    if (toggleButton) {
        toggleButton.addEventListener('click', () => {
            body.classList.toggle('dark-mode');
            if (body.classList.contains('dark-mode')) {
                localStorage.setItem('theme', 'dark');
                toggleButton.textContent = '☀️';
            } else {
                localStorage.setItem('theme', 'light');
                toggleButton.textContent = '🌙';
            }
        });
    }

    if (searchInput) {
        searchInput.addEventListener('input', (e) => {
            const searchTerm = e.target.value.toLowerCase().trim();

            cards.forEach(card => {
                let artistName = "";
                const h2 = card.querySelector('h2');
                const engraved = card.querySelector('.engraved-title');

                if (h2) artistName += h2.textContent.toLowerCase();
                if (engraved) artistName += " " + engraved.textContent.toLowerCase();

                if (artistName.includes(searchTerm) || searchTerm === "") {
                    card.style.display = 'block';
                } else {
                    card.style.display = 'none';
                }
            });
        });
    }
});