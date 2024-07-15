import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost/api', // バックエンドのベースURLを設定
});

export default api;
