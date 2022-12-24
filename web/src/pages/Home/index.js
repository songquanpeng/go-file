import React, { useEffect, useState } from 'react';
import { API, copy, showError, showSuccess } from '../../helpers';
import { useDropzone } from 'react-dropzone';
import { ITEMS_PER_PAGE } from '../../constants';
import { ReactComponent as DeleteIcon } from './delete.svg';
import { ReactComponent as QrCodeIcon } from './qr_code.svg';
import { ReactComponent as CopyIcon } from './copy.svg';
import { ReactComponent as PlayIcon } from './play.svg';
import { ReactComponent as DownloadIcon } from './download.svg';

const Home = () => {
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searching, setSearching] = useState(false);
  const [activePage, setActivePage] = useState(1);
  const [searchKeyword, setSearchKeyword] = useState('');
  const [searchLoading, setSearchLoading] = useState(false);
  const { acceptedFiles, getRootProps, getInputProps } = useDropzone();
  const [uploading, setUploading] = useState(false);
  const [uploadProgress, setUploadProgress] = useState('0');

  const loadFiles = async (startIdx) => {
    const res = await API.get(`/api/file/?p=${startIdx}`);
    const { success, message, data } = res.data;
    if (success) {
      if (startIdx === 0) {
        setFiles(data);
      } else {
        let newFiles = files;
        newFiles.push(...data);
        setFiles(newFiles);
      }
    } else {
      showError(message);
    }
    setLoading(false);
  };

  const onPaginationChange = (e, { activePage }) => {
    (async () => {
      if (activePage === Math.ceil(files.length / ITEMS_PER_PAGE) + 1) {
        // In this case we have to load more data and then append them.
        await loadFiles(activePage - 1);
      }
      setActivePage(activePage);
    })();
  };

  useEffect(() => {
    loadFiles(0)
      .then()
      .catch((reason) => {
        showError(reason);
      });
  }, []);

  const downloadFile = (link, filename) => {
    let linkElement = document.createElement('a');
    linkElement.download = filename;
    linkElement.href = '/upload/' + link;
    linkElement.click();
  };

  const copyLink = (link) => {
    let url = window.location.origin + '/upload/' + link;
    copy(url).then();
    showSuccess('链接已复制到剪贴板');
  };

  const deleteFile = async (id, idx) => {
    const res = await API.delete(`/api/file/${id}`);
    const { success, message } = res.data;
    if (success) {
      let newFiles = [...files];
      let realIdx = (activePage - 1) * ITEMS_PER_PAGE + idx;
      newFiles[realIdx].deleted = true;
      // newFiles.splice(idx, 1);
      setFiles(newFiles);
      showSuccess('文件已删除！');
    } else {
      showError(message);
    }
  };

  const searchFiles = async () => {
    if (searchKeyword === '') {
      // if keyword is blank, load files instead.
      await loadFiles(0);
      setActivePage(1);
      setSearching(false);
      return;
    }
    setSearchLoading(true);
    setSearching(true);
    const res = await API.get(`/api/file/search?keyword=${searchKeyword}`);
    const { success, message, data } = res.data;
    if (success) {
      setFiles(data);
      setActivePage(1);
      setSearchKeyword('');
    } else {
      showError(message);
    }
    setSearchLoading(false);
  };

  const handleKeywordChange = async (e) => {
    setSearchKeyword(e.target.value.trim());
  };

  const sortFile = (key) => {
    if (files.length === 0) return;
    setLoading(true);
    let sortedUsers = [...files];
    sortedUsers.sort((a, b) => {
      return ('' + a[key]).localeCompare(b[key]);
    });
    if (sortedUsers[0].id === files[0].id) {
      sortedUsers.reverse();
    }
    setFiles(sortedUsers);
    setLoading(false);
  };

  const uploadFiles = async () => {
    if (acceptedFiles.length === 0) return;
    setUploading(true);
    let formData = new FormData();
    for (let i = 0; i < acceptedFiles.length; i++) {
      formData.append('file', acceptedFiles[i]);
    }
    const res = await API.post(`/api/file`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      onUploadProgress: (e) => {
        let uploadProgress = ((e.loaded / e.total) * 100).toFixed(2);
        setUploadProgress(uploadProgress);
      },
    });
    const { success, message } = res.data;
    if (success) {
      showSuccess(`${acceptedFiles.length} 个文件上传成功！`);
    } else {
      showError(message);
    }
    setUploading(false);
    setUploadProgress('0');
    setSearchKeyword('');
    loadFiles(0).then();
    setActivePage(1);
  };

  useEffect(() => {
    uploadFiles().then();
  }, [acceptedFiles]);

  return (
    <>
      <div className='container'>
        <div className='normal-container'>
          <article
            id='messageToast'
            className='message is-danger'
            style={{ display: 'none' }}
          >
            <div className='message-body' id='messageToastText'></div>
          </article>
          <div
            className={'control' + (searchLoading ? ' is-loading' : '')}
            style={{ marginBottom: '16px' }}
          >
            <input
              className='input'
              type='text'
              placeholder='搜索文件的名称，上传者以及描述信息 ...'
              value={searchKeyword}
              onChange={handleKeywordChange}
              onKeyDown={(e) => {
                if (e.key === 'Enter') {
                  searchFiles().then();
                }
              }}
            />
          </div>
          <div className='box' id='fileUploadCard' style={{ display: 'none' }}>
            <article className='media'>
              <div className='media-content'>
                <div className='content'>
                  <div className='page-card-title' id='fileUploadTitle'></div>
                  <progress
                    className='progress is-success'
                    value='0'
                    max='100'
                    id='fileUploadProgress'
                  ></progress>
                  <div className='page-card-text' id='fileUploadDetail'></div>
                </div>
              </div>
            </article>
          </div>
          {files.length === 0 ? (
            <div className='box'>
              <div className='table-container'>
                <div className='box'>
                  <article className='media'>
                    <div className='media-content'>
                      <div className='content'>
                        <div
                          className='page-card-title'
                          style={{ color: '#AAAAAA' }}
                          id='imageUploadStatus'
                        >
                          {searching ? (
                            <>无匹配文件</>
                          ) : (
                            <>
                              当前无任何文件，请点击{' '}
                              <a
                                className='button is-light'
                                // onClick='showUploadModal()'
                              >
                                上传
                              </a>
                              按钮以上传文件.
                            </>
                          )}
                        </div>
                      </div>
                    </div>
                  </article>
                </div>
              </div>
            </div>
          ) : (
            <></>
          )}
          <div>
            {files
              .slice(
                (activePage - 1) * ITEMS_PER_PAGE,
                activePage * ITEMS_PER_PAGE
              )
              .map((file, idx) => {
                if (file.deleted) return <></>;
                return (
                  <div className='box file-box' id='file-{{$file.Id}}'>
                    <article className='media'>
                      <div className='media-content'>
                        <div className='content'>
                          <div className='page-card-title'>
                            <a
                              download='{{$file.Filename}}'
                              href={'/upload/' + file.link}
                            >
                              {file.filename}
                            </a>
                          </div>
                          <span className='tag is-light'>{file.uploader}</span>
                          <span className='tag is-light'>
                            {file.upload_time}
                          </span>
                          <span
                            className='tag is-light'
                            id='counter-{{$file.Id}}'
                          >
                            {file.download_counter} 次下载
                          </span>
                          <div className='page-card-text'>
                            {file.description ? file.description : '无描述信息'}
                          </div>
                          <div className='actions'>
                            <span
                              onClick={() => {
                                deleteFile(file.id, idx).then();
                              }}
                            >
                              <DeleteIcon />
                            </span>
                            <span onClick="showQRCode('/upload/{{$file.Link}}')">
                              <QrCodeIcon />
                            </span>
                            <span
                              onClick={() => {
                                copyLink(file.link);
                              }}
                            >
                              <CopyIcon />
                            </span>
                            <a target='_blank' href={'/upload/' + file.link}>
                              <PlayIcon />
                            </a>
                            <a
                              download={file.filename}
                              href={'/upload/' + file.link}
                              onClick={() => {}} // updateDownloadCounter('counter-{{$file.Id}}')
                            >
                              <DownloadIcon />
                            </a>
                          </div>
                        </div>
                      </div>
                    </article>
                  </div>
                );
              })}
          </div>
          <nav
            className='pagination is-centered'
            role='navigation'
            aria-label='pagination'
          >
            <a className='pagination-previous shadow' href='/?p={{.prev}}'>
              上一页
            </a>
            <a className='pagination-next shadow' href='/?p={{.next}}'>
              下一页
            </a>
          </nav>
        </div>
      </div>
    </>
  );
};

export default Home;
