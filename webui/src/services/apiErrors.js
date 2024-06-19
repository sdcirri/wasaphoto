
export const InternalServerError = new Error("internal server error");
export const BadIdsException = new Error("bad auth token or bad userID");
export const BlockedException = new Error("forbidden: user blocked you!");
export const UserNotFoundException = new Error("user not found");
export const PostNotFoundException = new Error("post not found");
export const BadAuthException = new Error("bad auth token");
export const BadFollowOperation = new Error("bad follow operation");
export const AccessDeniedException = new Error("access denied");
export const BadUploadException = new Error("image too big or text too big");
export const ImageTooBigException = new Error("image too big");
export const BadPostAuthException = new Error("cannot post as somebody else");
export const BadFeedException = new Error("trying to view somebody else's feed");
export const LikeImpersonationException = new Error("cannot like as somebody else");
export const UsernameAlreadyTakenException = new Error("username already taken");