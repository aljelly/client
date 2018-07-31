// @flow
import * as I from 'immutable'
import HiddenString from '../../util/hidden-string'

export type _State = {
  // TODO remove
  forgotPasswordError: ?Error,
  forgotPasswordSubmitting: boolean,
  forgotPasswordSuccess: boolean,
  justDeletedSelf: ?string,
  justRevokedSelf: ?string,
  // shared by all errors, we only ever want one error
  error: HiddenString,
  registerUserPassLoading: boolean,
}

export type State = I.RecordOf<_State>
