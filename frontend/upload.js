const serverUrl = 'http://localhost:8888';

function uploadFile() {
    const fileInput = document.getElementById('fileInput');
    const noteIdInput = document.getElementById('noteId');
    const file = fileInput.files[0];

    if (!file) {
        alert('Please select a file to upload.');
        return;
    }

    const formData = new FormData();
    formData.append('fileName', file.name);
    formData.append('noteId', Number(noteIdInput.value)); // 确保这是数字
    formData.append('fileContent', file); // 直接将文件作为fileContent发送

    // 获取当前时间戳（毫秒），如果后端需要秒为单位，请除以1000
    const timestamp = Date.now(); // 或者 Math.floor(Date.now() / 1000)
    formData.append('timestamp', timestamp.toString()); // 转换为字符串（尽管数字通常也可以）

    // 调用API
    fetch(serverUrl + '/note/upload', {
        method: 'POST',
        body: formData,
        // 注意：不要在这里设置'Content-Type'，浏览器会自动处理
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Success:', data);
            alert('File uploaded successfully!');
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error uploading file!');
        });
}