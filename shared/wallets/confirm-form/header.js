// @flow
import * as React from 'react'
import * as Kb from '../../common-adapters'
import * as Styles from '../../styles'

type HeaderProps = {|
  onBack: () => void,
  amount: string,
  assetType: string,
  assetConversion?: string,
  worthDescription?: string,
|}

const Header = (props: HeaderProps) => (
  <Kb.Box2 direction="horizontal" fullWidth={true} style={styles.header}>
    <Kb.Box2 direction="vertical" fullWidth={true} centerChildren={true} style={styles.headerContent}>
      <Kb.Icon
        type={
          Styles.isMobile
            ? 'icon-fancy-stellar-sending-mobile-149-129'
            : 'icon-fancy-stellar-sending-desktop-98-86'
        }
        style={Kb.iconCastPlatformStyles(styles.headerIcon)}
      />
      <Kb.Text selectable={true} type="BodyTiny" style={styles.headerText}>
        {(props.assetType === 'XLM' ? 'Sending' : 'Sending Lumens worth').toUpperCase()}
      </Kb.Text>
      <Kb.Text selectable={true} type="HeaderBigExtrabold" style={styles.headerText}>
        {props.assetConversion ? props.assetConversion : props.amount}
      </Kb.Text>
      {props.assetType === 'XLM' &&
        !!props.worthDescription && (
          <Kb.Text selectable={true} type="BodyTiny" style={styles.headerText}>
            {'(APPROXIMATELY ' + props.worthDescription + ')'}
          </Kb.Text>
        )}
    </Kb.Box2>
    <Kb.BackButton
      onClick={props.onBack}
      style={styles.backButton}
      iconColor={Styles.globalColors.white}
      textStyle={styles.backButtonText}
    />
  </Kb.Box2>
)

const styles = Styles.styleSheetCreate({
  header: Styles.platformStyles({
    common: {
      backgroundColor: Styles.globalColors.purple,
    },
    isElectron: {
      flex: 1,
      minHeight: 144,
    },
    isMobile: {
      flexGrow: 1,
      flexShrink: 1,
      flexBasis: 'auto',
      minHeight: 200,
    },
  }),
  headerContent: Styles.platformStyles({
    isElectron: {
      marginTop: -20,
    },
  }),
  headerText: {
    color: Styles.globalColors.white,
  },
  headerIcon: {
    marginBottom: Styles.globalMargins.small,
  },
  backButton: Styles.platformStyles({
    common: {
      position: 'absolute',
    },
    isElectron: {
      top: Styles.globalMargins.small,
      left: Styles.globalMargins.small,
    },
    isMobile: {
      top: Styles.globalMargins.tiny,
      left: Styles.globalMargins.tiny,
    },
  }),
  backButtonText: {
    color: Styles.globalColors.white,
  },
})

export default Header
