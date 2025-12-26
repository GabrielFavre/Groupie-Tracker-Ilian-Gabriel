document.addEventListener('DOMContentLoaded', () => {
    const toggleButton = document.getElementById('theme-toggle');
    const body = document.body;

    if (localStorage.getItem('theme') === 'dark') {
        body.classList.add('dark-mode');
        if(toggleButton) toggleButton.textContent = '☀️';
    }

    if (toggleButton) {
        toggleButton.addEventListener('click', () => {
            body.classList.toggle('dark-mode');

            // Sauvegarde le choix et change l'icône
            if (body.classList.contains('dark-mode')) {
                localStorage.setItem('theme', 'dark');
                toggleButton.textContent = '☀️';
            } else {
                localStorage.setItem('theme', 'light');
                toggleButton.textContent = '🌙';
            }
        });
    }
});