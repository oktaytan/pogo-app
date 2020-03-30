import { AXIOS } from '../../utils/settings';
import {
	FETCH_ALL_POSTS,
	SET_ALL_POSTS,
	GET_ALL_POSTS,
	FETCH_USERS_POSTS,
	SET_USERS_POSTS,
	GET_USERS_POSTS,
	FETCH_POST_BY_ID,
	SET_POST_BY_ID,
	GET_POST_BY_ID,
	FETCH_POST_BY_LIKES,
	SET_POST_BY_LIKES,
	GET_POST_BY_LIKES,
	POST_LIKE,
	ADD_NEW_POST,
	GET_NEW_POST_LOADING,
	DELETE_POST,
	FETCH_USER_LIKES,
	SET_USER_LIKES,
	GET_USER_LIKES,
	USER_LIKED,
	ADD_COMMENT,
	DELETE_COMMENT,
	FETCH_SEARCH_POSTS,
	SET_SEARCH_POSTS,
	GET_SEARCH_POSTS,
	FETCH_ALL_COMMENTS,
	SET_ALL_COMMENTS,
	GET_ALL_COMMENTS
} from '../keys';

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
	[GET_ALL_POSTS]: (state) => state.allPosts,
	[GET_USERS_POSTS]: (state) => state.usersPosts,
	[GET_POST_BY_ID]: (state) => state.postById,
	[GET_USER_LIKES]: (state) => state.userLikes,
	[GET_POST_BY_LIKES]: (state) => state.mostLikedPosts,
	[GET_SEARCH_POSTS]: (state) => state.searchedPosts,
	[GET_ALL_COMMENTS]: (state) => state.allComments,
	[GET_NEW_POST_LOADING]: (state) => state.newPostLoading
};

const actions = {
	[FETCH_ALL_POSTS]: ({ commit }) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/posts')
				.then((res) => {
					commit('SET_ALL_POSTS', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[FETCH_USERS_POSTS]: ({ commit }, username) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/${username}/posts`)
				.then((res) => {
					commit('SET_USERS_POSTS', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[FETCH_ALL_COMMENTS]: ({ commit }, username) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/comments')
				.then((res) => {
					commit('SET_ALL_COMMENTS', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[FETCH_POST_BY_ID]: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/posts/${id}`)
				.then((res) => {
					commit('SET_POST_BY_ID', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[FETCH_POST_BY_LIKES]: ({ commit }) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/posts/liked')
				.then((res) => {
					commit('SET_POST_BY_LIKES', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[FETCH_SEARCH_POSTS]: ({ commit }, search) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/posts/search', JSON.stringify(search))
				.then((res) => {
					commit('SET_SEARCH_POSTS', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[ADD_NEW_POST]: ({ commit }, post) => {
		state.newPostLoading = true;
		return new Promise((resolve, reject) => {
			AXIOS.post('/posts', JSON.stringify(post))
				.then((resData) => {
					AXIOS.get('/posts')
						.then((res) => {
							commit('SET_ALL_POSTS', res.data);
							resolve(res.data);
							state.newPostLoading = false;
						})
						.catch((err) => reject(err));
					resolve(resData.data);
				})
				.catch((err) => reject(err));
		});
	},

	[DELETE_POST]: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.delete(`/posts/${id}`)
				.then((resData) => {
					AXIOS.get('/posts')
						.then((res) => {
							commit('SET_ALL_POSTS', res.data);
							resolve(res.data);
						})
						.catch((err) => reject(err));
					resolve(resData.data);
				})
				.catch((err) => reject(err));
		});
	},

	[POST_LIKE]: ({ commit }, payload) => {
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

	[FETCH_USER_LIKES]: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/likes/${id}`)
				.then((res) => {
					commit('SET_USER_LIKES', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[USER_LIKED]: ({ commit }, payload) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/likes', JSON.stringify(payload))
				.then((res) => {
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	[ADD_COMMENT]: ({ commit }, comment) => {
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

	[DELETE_COMMENT]: ({ commit }, comment_id) => {
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
	[SET_ALL_POSTS]: (state, posts) => (state.allPosts = posts),
	[SET_USERS_POSTS]: (state, posts) => (state.usersPosts = posts),
	[SET_POST_BY_ID]: (state, post) => (state.postById = post),
	[SET_USER_LIKES]: (state, likes) => (state.userLikes = likes),
	[SET_POST_BY_LIKES]: (state, posts) => (state.mostLikedPosts = posts),
	[SET_SEARCH_POSTS]: (state, posts) => (state.searchedPosts = posts),
	[SET_ALL_COMMENTS]: (state, comments) => (state.allComments = comments)
};

export default {
	state,
	getters,
	actions,
	mutations
};
