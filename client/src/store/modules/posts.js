import { AXIOS } from '../../utils/settings';
import * as _ from '../keys';

const state = {
	allPosts: [],
	usersPosts: [],
	postById: null,
	userLikes: [],
	mostLikedPosts: [],
	searchedPosts: [],
	allComments: [],
	newPostLoading: false
};

const getters = {
	[_.GET_ALL_POSTS]: (state) => state.allPosts,
	[_.GET_USERS_POSTS]: (state) => state.usersPosts,
	[_.GET_POST_BY_ID]: (state) => state.postById,
	[_.GET_USER_LIKES]: (state) => state.userLikes,
	[_.GET_POST_BY_LIKES]: (state) => state.mostLikedPosts,
	[_.GET_SEARCH_POSTS]: (state) => state.searchedPosts,
	[_.GET_ALL_COMMENTS]: (state) => state.allComments,
	[_.GET_NEW_POST_LOADING]: (state) => state.newPostLoading
};

const actions = {
	// Tüm gönderileri getirme
	[_.FETCH_ALL_POSTS]: ({ commit }) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/posts')
				.then((res) => {
					if (res.data) {
						commit(_.SET_ALL_POSTS, res.data);
						resolve(res.data);
					} else {
						const data = {
							emty: true,
							message: 'Henüz hiç gönderi yok!'
						};
						resolve(data);
					}
				})
				.catch((err) => reject(err));
		});
	},

	// Kullanıcının kendi gönderilerini getirme
	[_.FETCH_USERS_POSTS]: ({ commit }, username) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/${username}/posts`)
				.then((res) => {
					commit(_.SET_USERS_POSTS, res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Tüm yorumları getirme
	[_.FETCH_ALL_COMMENTS]: ({ commit }) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/comments')
				.then((res) => {
					commit(_.SET_ALL_COMMENTS, res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Seçili gönderiyi getirme
	[_.FETCH_POST_BY_ID]: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/posts/${id}`)
				.then((res) => {
					commit(_.SET_POST_BY_ID, res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Beğeni sayısı sıralamasına göre gönerileri getirme
	[_.FETCH_POST_BY_LIKES]: ({ commit }) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/posts/liked')
				.then((res) => {
					commit(_.SET_POST_BY_LIKES, res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Arama sonucuna göre gönderileri getirme
	[_.FETCH_SEARCH_POSTS]: ({ commit }, search) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/posts/search', JSON.stringify(search))
				.then((res) => {
					commit(_.SET_SEARCH_POSTS, res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Yeni gönderi ekleme
	[_.ADD_NEW_POST]: ({ commit }, post) => {
		state.newPostLoading = true;
		return new Promise((resolve, reject) => {
			AXIOS.post('/posts', JSON.stringify(post))
				.then((resData) => {
					AXIOS.get('/posts')
						.then((res) => {
							commit(_.SET_ALL_POSTS, res.data);
							resolve(res.data);
							state.newPostLoading = false;
						})
						.catch((err) => reject(err));
					resolve(resData.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Gönderi silme
	[_.DELETE_POST]: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.delete(`/posts/${id}`)
				.then((resData) => {
					AXIOS.get('/posts')
						.then((res) => {
							commit(_.SET_ALL_POSTS, res.data);
							resolve(res.data);
						})
						.catch((err) => reject(err));
					resolve(resData.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Gönderi beğenme
	[_.POST_LIKE]: ({ commit }, payload) => {
		let newLikes = {
			likes: payload.likes
		};
		return new Promise((resolve, reject) => {
			AXIOS.put(`/posts/${payload.id}`, JSON.stringify(newLikes))
				.then((res) => {
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Kullanıcının yaptığı beğenileri getirme
	[_.FETCH_USER_LIKES]: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/likes/${id}`)
				.then((res) => {
					commit(_.SET_USER_LIKES, res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Kullanıcı gönderiyi beğendi ve ya beğeniyi geri aldı durumu
	[_.USER_LIKED]: ({ commit }, payload) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/likes', JSON.stringify(payload))
				.then((res) => {
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	// Yorum ekleme
	[_.ADD_COMMENT]: ({ commit }, comment) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/comments', JSON.stringify(comment))
				.then((res) => {
					if (res.data.error) {
						reject(res.data);
					} else {
						resolve(res.data);
					}
				})
				.catch((err) => reject(err));
		});
	},

	// Yorum silme
	[_.DELETE_COMMENT]: ({ commit }, comment_id) => {
		return new Promise((resolve, reject) => {
			AXIOS.delete(`/comments/${comment_id}`)
				.then((res) => {
					if (res.data.error) {
						reject(res.data);
					} else {
						resolve(res.data);
					}
				})
				.catch((err) => reject(err));
		});
	}
};

const mutations = {
	[_.SET_ALL_POSTS]: (state, posts) => (state.allPosts = posts),
	[_.SET_USERS_POSTS]: (state, posts) => (state.usersPosts = posts),
	[_.SET_POST_BY_ID]: (state, post) => (state.postById = post),
	[_.SET_USER_LIKES]: (state, likes) => (state.userLikes = likes),
	[_.SET_POST_BY_LIKES]: (state, posts) => (state.mostLikedPosts = posts),
	[_.SET_SEARCH_POSTS]: (state, posts) => (state.searchedPosts = posts),
	[_.SET_ALL_COMMENTS]: (state, comments) => (state.allComments = comments)
};

export default {
	state,
	getters,
	actions,
	mutations
};
