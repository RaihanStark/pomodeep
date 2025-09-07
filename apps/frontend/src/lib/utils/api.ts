const API_BASE_URL = 'http://localhost:8080/api/v1';

export interface ApiError {
    error: string;
}

export interface SignUpRequest {
    email: string;
    password: string;
}

export interface SignUpResponse {
    user: {
        id: number;
        email: string;
        created_at: string;
        updated_at: string;
    };
}

export interface SignInRequest {
    email: string;
    password: string;
}

export interface SignInResponse {
    user: {
        id: number;
        email: string;
        created_at: string;
        updated_at: string;
    };
    token: string;
}

export async function signUp(data: SignUpRequest): Promise<SignUpResponse> {
    const response = await fetch(`${API_BASE_URL}/auth/signup`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        const errorData: ApiError = await response.json();
        throw new Error(errorData.error || 'Sign up failed');
    }

    return response.json();
}

export async function signIn(data: SignInRequest): Promise<SignInResponse> {
    const response = await fetch(`${API_BASE_URL}/auth/signin`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        const errorData: ApiError = await response.json();
        throw new Error(errorData.error || 'Sign in failed');
    }

    return response.json();
}

// Token management utilities
export function setAuthToken(token: string) {
    if (typeof window !== 'undefined') {
        localStorage.setItem('auth_token', token);
    }
}

export function getAuthToken(): string | null {
    if (typeof window !== 'undefined') {
        return localStorage.getItem('auth_token');
    }
    return null;
}

export function removeAuthToken() {
    if (typeof window !== 'undefined') {
        localStorage.removeItem('auth_token');
    }
}

export function isAuthenticated(): boolean {
    return getAuthToken() !== null;
}
