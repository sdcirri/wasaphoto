
export const InternalServerError = new Error("Internal server error");
export const BadIdsException = new Error("Bad auth token or bad userID");
export const BlockedException = new Error("Forbidden: user blocked you!");
export const UserNotFoundException = new Error("User not found");
export const BadAuthException = new Error("Bad auth token");
export const BadFollowOperation = new Error("Bad follow operation");
export const AccessDeniedException = new Error("Access denied");
export const BadUploadException = new Error("Image too big or text too big");
export const BadPostAuthException = new Error("Cannot post as somebody else");
