import React, { useEffect, useState } from 'react';

const Footer = () => {
  const [Footer, setFooter] = useState('');
  useEffect(() => {
    let savedFooter = localStorage.getItem('footer_html');
    if (!savedFooter) savedFooter = '';
    setFooter(savedFooter);
  });

  return (
    <footer className='footer' style={{ backgroundColor: 'white' }}>
      <div className='content has-text-centered'>
        {Footer === '' ? (
          <p>
            <a href='https://github.com/songquanpeng/go-file' target='_blank'>
              Go File {process.env.REACT_APP_VERSION}
            </a>
            由 <a href='https://github.com/songquanpeng'>JustSong</a> 构建，
            源代码遵循
            <a href='https://opensource.org/licenses/mit-license.php'> MIT </a>
            协议
          </p>
        ) : (
          <p dangerouslySetInnerHTML={{ __html: Footer }}></p>
        )}
      </div>
    </footer>
  );
};

export default Footer;
