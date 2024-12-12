
const jsonFile = 'rscr/scores.json';


const scoreboard = document.getElementById('scoreboard');


fetch('rscr/scores.json')
  .then(response => {
    if (!response.ok) {
      console.error('Erreur de chargement du fichier JSON :', response.statusText);
      throw new Error('Erreur de chargement des données JSON');
    }
    return response.json(); 
  })
  .then(data => {
    console.log('Données chargées avec succès :', data);
    scoreboard.innerHTML = '';

    
    data.forEach((player, index) => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${index + 1}</td>
        <td>${player.username}</td>
        <td>${player.score}</td>
      `;
      scoreboard.appendChild(row);
    });
  })
  .catch(error => {
    console.error('Erreur lors du chargement des données JSON :', error)
    scoreboard.innerHTML = '<tr><td colspan="3">Impossible de charger les scores.</td></tr>';
  });
