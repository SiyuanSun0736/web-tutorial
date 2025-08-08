document.addEventListener('DOMContentLoaded', function() {
    // 获取主页标题元素
    const title = document.querySelector('header h1');
    
    // 添加点击事件，点击标题时更改颜色
    if (title) {
        title.addEventListener('click', function() {
            const colors = ['#3498db', '#e74c3c', '#2ecc71', '#f39c12', '#9b59b6'];
            const randomColor = colors[Math.floor(Math.random() * colors.length)];
            this.style.color = randomColor;
        });
    }
    
    // 添加内容到主页
    const main = document.querySelector('main');
    if (main) {
        // 创建一个新段落
        const newParagraph = document.createElement('p');
        newParagraph.textContent = '点击上方标题可以改变它的颜色！';
        newParagraph.className = 'home-hint';
        
        // 添加到主内容区域
        main.appendChild(newParagraph);
        
        // 添加样式
        newParagraph.style.fontSize = '0.9em';
        newParagraph.style.fontStyle = 'italic';
        newParagraph.style.color = '#666';
    }
    
    console.log('home.js 已加载完成');
});