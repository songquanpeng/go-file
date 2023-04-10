import React from 'react';
import { Segment, Dimmer, Loader } from 'semantic-ui-react';

const Loading = ({ prompt: name = 'page' }) => {
  return (
    <div className="loading">
      <div className="spinner"></div>
    </div>
  );
};

export default Loading;
