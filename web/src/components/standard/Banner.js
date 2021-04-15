import styled from 'styled-components';
import { Colors } from '../../utils/Colors';

const Banner = styled.div`
  width: 100%;
  padding: 3px 8px;
`;

export const SuccessBanner = styled(Banner)`
  padding: 3px 8px;
  background-color: ${Colors.Green};
`;

export const ErrorBanner = styled(Banner)`
  padding: 3px 8px;
  background-color: ${Colors.Yellow};
`;
