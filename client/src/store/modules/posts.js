import { AXIOS } from '../../utils/settings';
import {
	FETCH_ALL_POSTS,
	SET_ALL_POSTS,
	GET_ALL_POSTS,
	FETCH_POST_BY_ID,
	SET_POST_BY_ID,
	GET_POST_BY_ID,
	POST_LIKE,
	ADD_COMMENT,
	DELETE_COMMENT
} from '../keys';

const state = {
	allPosts: [],
	postById: null,
	
};

const getters = {
	GET_ALL_POSTS: (state) => state.posts,
	GET_POST_BY_ID: (state) => state.postById
};

const actions = {
	FETCH_ALL_POSTS: ({ commit }) => {
		return new Promise((resolve, reject) => {
			AXIOS.get('/posts')
				.then((res) => {
					commit('SET_ALL_POSTS', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	FETCH_POST_BY_ID: ({ commit }, id) => {
		return new Promise((resolve, reject) => {
			AXIOS.get(`/posts/${id}`)
				.then((res) => {
					commit('SET_POST_BY_ID', res.data);
					resolve(res.data);
				})
				.catch((err) => reject(err));
		});
	},

	POST_LIKE: ({ commit }, payload) => {
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

	ADD_COMMENT: ({ commit }, comment) => {
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

	DELETE_COMMENT: ({ commit }, comment_id) => {
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
	SET_ALL_POSTS: (state, posts) => (state.posts = posts),
	SET_POST_BY_ID: (state, post) => (state.postById = post)
};

export default {
	state,
	getters,
	actions,
	mutations
};
