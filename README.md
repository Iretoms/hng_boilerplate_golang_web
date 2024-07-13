# API Documentation

## Overview

This document provides an overview of the API endpoints available in the system, including authentication, user management, payments, and more.

## Base URL

https://www.example.com/api/v1

## Table of Contents

- [Authentication](#authentication)
  - [Register](#register)
  - [Login](#login)
  - [Social Authentication](#social-authentication)
  - [Magic Link Authentication](#magic-link-authentication)
  - [Change Password](#change-password)

## Authentication 
 **Description:** Endpoints for user registration and authentication.

**Endpoint:** `POST /auth/register` **Description:** Register a new user.

**Endpoint:** `POST /auth/login` **Description:** Login a new user.

**Endpoint:** `POST /auth/social` **Description:** Authenticate a user using a social provider.

**Endpoint:** `POST /auth/magic-link` **Description:** Request a magic link for authentication.

**Endpoint:** `GET /auth/magic-link/verify` **Description:** Verify the magic link token.

**Endpoint:** `POST /auth/magic-link/password-reset` **Description:** Request a password reset link.

**Endpoint:** `POST /auth/magic-link/password-reset/verify` **Description:** Verify the password reset token and set a new password.

**Endpoint:** `PUT /auth/change-password` **Description:** Verify the password reset token and set a new password.


## Invite
 **Description:** API for user invite management.

 **Endpoint:** `POST /invite/generate` **Description:** Generates an invitation token.

 **Endpoint:** `POST /invite/verify` **Description:** Verify.


 ## Notifications
 **Description:** Notification operations.

 **Endpoint:** `POST /notifications/{user_id}` **Description:** Get notifications by user Id.


## Pages
 **Description:** Pages available.

 **Endpoint:** `POST /contact` **Description:** Submit a contact form.

 **Endpoint:** `GET /contacts/messages` **Description:** Retrieve all contact messages (Admin only).


 ## Payment
 **Description:** Payment operations.

 **Endpoint:** `GET /payments` **Description:** List all payments.

 **Endpoint:** `POST /payments` **Description:** Create a new payment.

 **Endpoint:** `GET /payments/{paymentID}` **Description:** Get payment by ID.

 **Endpoint:** `PUT /payments/{paymentID}` **Description:** Update an existing payment.

 **Endpoint:** `POST /payments/stripe` **Description:** Create a new payment via Stripe.

 **Endpoint:** `POST /payments/flutterwave` **Description:** Create a new payment via Flutterwave.

 **Endpoint:** `POST /payments/lemonsqueezy` **Description:** Create a new payment via Lemonsqueezy.
