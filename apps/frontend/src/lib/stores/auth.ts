import { writable } from 'svelte/store';
import { getAuthToken, removeAuthToken } from '$lib/utils/api';

export interface User {
    id: number;
    email: string;
    created_at: string;
    updated_at: string;
}

export interface AuthState {
    isAuthenticated: boolean;
    user: User | null;
    token: string | null;
}

// Initialize auth state from localStorage (only in browser)
function createAuthStore() {
    const token = typeof window !== 'undefined' ? getAuthToken() : null;
    const initialState: AuthState = {
        isAuthenticated: !!token,
        user: null, // We'll populate this when needed
        token
    };

    const { subscribe, set, update } = writable<AuthState>(initialState);

    return {
        subscribe,
        login: (user: User, token: string) => {
            update(state => ({
                ...state,
                isAuthenticated: true,
                user,
                token
            }));
        },
        logout: () => {
            removeAuthToken();
            set({
                isAuthenticated: false,
                user: null,
                token: null
            });
        },
        setUser: (user: User) => {
            update(state => ({
                ...state,
                user
            }));
        }
    };
}

export const authStore = createAuthStore();
