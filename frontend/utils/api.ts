import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost/api', // バックエンドのベースURLを設定
  withCredentials: true, // クッキーを含むリクエストを送信
});

export default api;
