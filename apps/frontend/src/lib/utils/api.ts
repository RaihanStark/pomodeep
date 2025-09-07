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
