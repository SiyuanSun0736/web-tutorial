document.addEventListener('DOMContentLoaded', function() {
    // 获取联系页面标题元素
    const title = document.querySelector('header h1');
    
    // 添加点击事件，点击标题时更改颜色
    if (title) {
        title.addEventListener('click', function() {
            const colors = ['#e74c3c', '#f39c12', '#1abc9c', '#8e44ad', '#2c3e50'];
            const randomColor = colors[Math.floor(Math.random() * colors.length)];
            this.style.color = randomColor;
        });
    }
    
    // 表单验证和提交处理
    const contactForm = document.querySelector('.contact-form form');
    if (contactForm) {
        contactForm.addEventListener('submit', function(e) {
            e.preventDefault();
            
            const name = document.getElementById('name').value.trim();
            const email = document.getElementById('email').value.trim();
            const message = document.getElementById('message').value.trim();
            
            if (!name || !email || !message) {
                alert('请填写所有必填字段！');
                return;
            }
            
            if (!isValidEmail(email)) {
                alert('请输入有效的邮箱地址！');
                return;
            }
            
            // 模拟提交成功
            alert('消息发送成功！我们会尽快回复您。');
            this.reset();
        });
    }
    
    // 为图片添加悬停效果
    const contactImage = document.querySelector('.contact-content img');
    if (contactImage) {
        contactImage.addEventListener('mouseenter', function() {
            this.style.opacity = '0.8';
            this.style.transition = 'opacity 0.3s ease';
        });
        
        contactImage.addEventListener('mouseleave', function() {
            this.style.opacity = '1';
        });
    }
    
    // 邮箱格式验证函数
    function isValidEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }
    
    // 添加动态内容
    const main = document.querySelector('main .contact-content');
    if (main) {
        const newParagraph = document.createElement('p');
        newParagraph.textContent = '填写下方表单与我们联系，或点击标题改变颜色！';
        newParagraph.className = 'contact-hint';
        newParagraph.style.fontSize = '0.9em';
        newParagraph.style.fontStyle = 'italic';
        newParagraph.style.color = '#666';
        newParagraph.style.textAlign = 'center';
        main.insertBefore(newParagraph, main.querySelector('.contact-form'));
    }
    
    console.log('contact.js 已加载完成');
});