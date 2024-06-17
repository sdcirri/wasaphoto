import axios from "axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

instance.interceptors.response.use(
	function (response) {		// 2xx
		return response;
	},
	function (error) {
		if (error.response &&
			error.response.status >= 400 &&
			error.response.status <= 404)
			return Promise.resolve(error.response);
		Promise.reject(error);
	}
);

export default instance;
