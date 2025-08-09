document.addEventListener('DOMContentLoaded', function() {
    // 获取关于页面标题元素
    const title = document.querySelector('header h1');
    
    // 添加点击事件，点击标题时更改颜色
    if (title) {
        title.addEventListener('click', function() {
            const colors = ['#3498db', '#2ecc71', '#e67e22', '#9b59b6', '#34495e'];
            const randomColor = colors[Math.floor(Math.random() * colors.length)];
            this.style.color = randomColor;
        });
    }
    
    // 为图片添加点击效果
    const aboutImage = document.querySelector('.about-content img');
    if (aboutImage) {
        aboutImage.addEventListener('click', function() {
            this.style.transform = this.style.transform === 'scale(1.1)' ? 'scale(1)' : 'scale(1.1)';
            this.style.transition = 'transform 0.3s ease';
        });
    }
    
    // 添加动态内容
    const main = document.querySelector('main .about-content');
    if (main) {
        const newParagraph = document.createElement('p');
        newParagraph.textContent = '点击标题改变颜色，点击图片放大缩小！';
        newParagraph.className = 'about-hint';
        newParagraph.style.fontSize = '0.9em';
        newParagraph.style.fontStyle = 'italic';
        newParagraph.style.color = '#666';
        main.appendChild(newParagraph);
    }
    
    console.log('about.js 已加载完成');
});