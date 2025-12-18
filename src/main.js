import './style.css';

const shutdownBtn = document.getElementById('shutdownBtn');
const closeBtn = document.getElementById('closeBtn');
const statusMessage = document.getElementById('status-message');

const ip = "http://192.168.122.13:3000";

function updateStatus(message, isError = false) {
  statusMessage.textContent = message;
  statusMessage.className = isError ? 'error' : 'success';
  setTimeout(() => {
    statusMessage.textContent = '';
    statusMessage.className = '';
  }, 3000);
}

if (shutdownBtn) {
  shutdownBtn.addEventListener('click', () => {
    updateStatus('Sending shutdown command...');
    fetch(`${ip}/shutdown`)
      .then(response => {
        if (response.ok) {
          updateStatus('Shutdown command sent successfully!');
        } else {
          updateStatus('Failed to send shutdown command.', true);
        }
      })
      .catch(error => {
        console.error('Error:', error);
        updateStatus('An error occurred while sending shutdown command.', true);
      });
  });
}

if (closeBtn) {
  closeBtn.addEventListener('click', () => {
    updateStatus('Sending close command...');
    fetch(`${ip}/close`)
      .then(response => {
        if (response.ok) {
          updateStatus('Close command sent successfully!');
        } else {
          updateStatus('Failed to send close command.', true);
        }
      })
      .catch(error => {
        console.error('Error:', error);
        updateStatus('An error occurred while sending close command.', true);
      });
  });
}

