import styled from 'styled-components';
import { Colors } from '../../utils/Colors';

const Banner = styled.div`
  width: 100%;
  padding: 5px 10px;
`;

export const SuccessBanner = styled(Banner)`
  background-color: ${Colors.Green};
`;

export const ErrorBanner = styled(Banner)`
  background-color: ${Colors.Yellow};
`;
