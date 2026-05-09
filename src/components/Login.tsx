import React, { useState } from 'react';

export function LoginButton() {
  const [hovered, setHovered] = useState(false);
  return (
    <button
      type="button"
      title="Sign in to your account"
      aria-label="Sign in to your account"
      data-testid="login-button"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      style={{
        display: 'inline-flex',
        alignItems: 'center',
        gap: 8,
        padding: '10px 18px',
        borderRadius: 8,
        border: 'none',
        cursor: 'pointer',
        fontSize: 16,
        fontWeight: 600,
        color: 'white',
        background: hovered ? '#1d4ed8' : '#2563eb',
        transition: 'background 0.15s ease',
      }}
    >
      <svg
        data-testid="login-button-icon"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        aria-hidden="true"
      >
        <rect x="5" y="11" width="14" height="9" rx="2"></rect>
        <path d="M8 11V8a4 4 0 0 1 8 0v3"></path>
      </svg>
      Sign in
    </button>
  );
}
